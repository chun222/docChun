/*
 * @Date: 2022-02-25 08:52:05
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-05-08 10:40:46
 * @FilePath: \server\system\core\request\request.go
 */
package request

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"chunDoc/system/core/trans"
	"chunDoc/system/util/convert"
	"chunDoc/system/util/str"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

//获取绑定参数并验证
func Binding(obj interface{}, c *gin.Context) (error, string) {
	err := c.ShouldBindJSON(&obj)

	// t := reflect.TypeOf(obj)
	//
	// if t.Kind() == reflect.Ptr {
	// 	t = t.Elem()
	// }
	// if t.Kind() != reflect.Struct {
	// 	fieldNum := t.NumField()
	// 	for i := 0; i < fieldNum; i++ {
	// 		tagName := t.Field(i).Name
	// 		tagTag := t.Field(i).Tag
	// 		fmt.Println(tagName, tagTag)
	// 	}
	// }

	msg := ""
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			msg = err.Error()
		} else {
			mapS := errs.Translate(trans.Trans)

			for _, v := range mapS {
				msg += v
				break //只返回一项报错，避免翻译问题
			}
		}
	}
	return err, msg
}

//生成查询语句
func GenWhere(obj interface{}, db *gorm.DB) {
	// if obj.LoginName != "" {
	// 	db = db.Where("user_name LIKE ?", "%"+obj.LoginName+"%")
	// }
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	//指针
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if t.Kind() == reflect.Struct {
		fieldNum := t.NumField()

		for i := 0; i < fieldNum; i++ {
			//tagName := t.Field(i).Name

			tagTag := t.Field(i).Tag

			tagValue := v.Field(i)
			fieldname := tagTag.Get("json")
			searchtype := tagTag.Get("search")

			if searchtype != "" && fieldname != "" {
				//前端传回来只接收数字和字符串
				if tagValue.Kind() == reflect.String {
					if tagValue.Interface().(string) != "" {
						sqlstring(fieldname, tagValue.Interface().(string), searchtype, db)
					}
				} else if tagValue.Kind() == reflect.Slice {
					if len(tagValue.Interface().([]time.Time)) > 1 {
						sqlstring(fieldname, tagValue.Interface(), searchtype, db)
					}
				} else {
					if convert.String(tagValue.Interface()) != "0" {
						sqlstring(fieldname, tagValue.Interface(), searchtype, db)
					}

				}
			}

		}
	}
}

func sqlstring(name string, v interface{}, f string, db *gorm.DB) {

	stmt := &gorm.Statement{DB: db}
	stmt.Parse(&db.Statement.Model)
	maintablename := stmt.Schema.Table

	if str.Contians(name, ".") {
		names := strings.Split(name, ".")
		//转驼峰以连表的一致
		name = fmt.Sprintf("`%s`.`%s`", str.Case2Camel(names[0]), names[len(names)-1])
	} else {
		name = fmt.Sprintf("`%s`.`%s`", maintablename, name)
	}

	if f == "btw" || f == "between" {
		//逗号分隔
		// db.Where(name+" "+" BETWEEN  ? and ?", str.BeforeLast(v.(string), ","), str.AfterLast(v.(string), ","))

		db.Where(name+" "+" BETWEEN  ? and ?", v.([]time.Time)[0], v.([]time.Time)[1])

	} else if str.Contians(f, "like") {
		db.Where(name+" "+f+" ?", "%"+v.(string)+"%")
	} else {
		db.Where(name+" "+f+" ?", v)
	}

	//fmt.Println(name+" "+f+" ?", v, "================")

	// switch f {
	// case "like","LIKE":
	// 	db.Where(name+" "+" LIKE ?", v)

	// case "=":
	// 	db.Where(name+" "+" = ?", v)

	// case "<>":
	// 	db.Where(name+" "+" <> ?", v)

	// }

}

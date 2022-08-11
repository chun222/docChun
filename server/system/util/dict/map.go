/*
 * @Date: 2022-02-22 11:29:54
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-02-28 16:49:35
 * @FilePath: \gin-antd\server\app\util\dict\map.go
 */
package dict

import "reflect"

func Delete(src map[string]interface{}, args ...string) {
	for _, v := range args {
		delete(src, v)
	}
}


//binding type interface 要修改的结构体
//value type interace 有数据的结构体
func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
} 

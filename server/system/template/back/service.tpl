/* 
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 自动生成-{{.StructName}}
 */

package Service

import (
	"chunDoc/system/core/db"
	"chunDoc/system/core/request"
	"chunDoc/system/model/CommonModel"
	AppDbModel "chunDoc/app/model/DbModel"  
	AppRequestModel "chunDoc/app/model/RequestModel"
)

type {{.StructName}}Service struct{}

//新增
func ({{.StructName}}Service *{{.StructName}}Service) Add(info *AppDbModel.{{.StructName}}) (err error) {
	result := db.Instance().Create(info)
	return result.Error
}

//修改
func ({{.StructName}}Service *{{.StructName}}Service) Edit(info *AppDbModel.{{.StructName}}) (err error) {
	var data AppDbModel.{{.StructName}}
	if err := db.Instance().First(&data, info.Id).Error; err != nil {
		return err
	}
	//忽略字段
	return db.Instance().Omit("CreatedAt").Save(info).Error
}

//分页列表
func ({{.StructName}}Service *{{.StructName}}Service) List(info AppRequestModel.{{.StructName}}) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := db.Instance().Model(&AppDbModel.{{.StructName}}{})
	//动态查询条件
	request.GenWhere(info, db)
	var dataList []AppDbModel.{{.StructName}}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&dataList).Error
	return err, dataList, total
}

//删除
func ({{.StructName}}Service *{{.StructName}}Service) Remove(info CommonModel.Ids) (err error) {
	result := db.Instance().Delete(&AppDbModel.{{.StructName}}{}, info.Ids)
	return result.Error
}
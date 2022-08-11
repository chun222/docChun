/*
 * @Date: 2022-02-14 15:13:57
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-05-08 10:30:07
 * @FilePath: \server\system\model\CommonModel\common.go
 */
package CommonModel

import (
	"time"
)

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`                 // 主键ID
	CreatedAt time.Time `json:"create_time" gorm:"column:create_time;comment:创建时间"` // 创建时间
	UpdatedAt time.Time `json:"update_time" gorm:"column:update_time;comment:更新时间"` // 更新时间
}

type OperaterModel struct {
	OperaterId   uint   `json:"operater_id" gorm:"column:operater_id"`
	OperaterUser string `json:"operater_user" gorm:"column:operater_user;size:200"`
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

//响应结果
const (
	SUCCESS      = 0   // 成功
	ERROR        = 500 //错误
	UNAUTHORIZED = 403 //无权限
	FAIL         = -1  //失败
)

// 请求页码
type PageInfo struct {
	Page     int `json:"page" form:"page" binding:"required"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize" binding:"required"` // 每页大小
}
type Id struct {
	Id uint `json:"id" form:"id" binding:"required" search:"="`
}

type Ids struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

type Names struct {
	Names []string `json:"names" form:"names" binding:"required"`
}

type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}

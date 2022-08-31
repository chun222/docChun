/*
 * @Date: 2022-04-12 09:11:39
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-08-31 17:46:15
 * @FilePath: \server\system\model\RequestModel\tools.go
 */
package RequestModel

import "mime/multipart"

type Search struct {
	Lang    string `json:"lang" binding:"required"`
	Version string `json:"version" binding:"required"`
	Keyword string `json:"keyword" binding:"required"`
}

type InitData struct {
	Lang    string `json:"lang" binding:"required"`
	Version string `json:"version" binding:"required"`
	Project string `json:"project" binding:"required"`
}

type Path struct {
	InitData
	Page    string `json:"page" binding:"required"`
	Content string
}

type FileUpload struct {
	File    multipart.FileHeader `form:"file" json:"file"`
	Path    []string             `form:"path[]" json:"path"`
	Dirname string               `json:"dirname"` //新建文件夹用到
}

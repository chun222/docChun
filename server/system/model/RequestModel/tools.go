/*
 * @Date: 2022-04-12 09:11:39
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-08-15 17:35:02
 * @FilePath: \server\system\model\RequestModel\tools.go
 */
package RequestModel

import "mime/multipart"

type Search struct {
	LangName string `json:"langname" binding:"required"`
	Keyword  string `json:"keyword" binding:"required"`
}

type LangName struct {
	LangName string `json:"langname" binding:"required"`
}

type Path struct {
	Path string `json:"path" binding:"required"`
}

type FileUpload struct {
	File    multipart.FileHeader `form:"file" json:"file"`
	Path    []string             `form:"path[]" json:"path"`
	Dirname string               `json:"dirname"` //新建文件夹用到
}

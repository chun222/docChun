/*
 * @Date: 2022-04-12 09:11:39
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-08-11 17:41:48
 * @FilePath: \server\system\model\RequestModel\tools.go
 */
package RequestModel

import "mime/multipart"

type LangName struct {
	LangName string `json:"langname" binding:"required"`
}

type FileUpload struct {
	File    multipart.FileHeader `form:"file" json:"file"`
	Path    []string             `form:"path[]" json:"path"`
	Dirname string               `json:"dirname"` //新建文件夹用到
}

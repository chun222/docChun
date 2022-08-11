/*
 * @Date: 2022-04-12 09:06:05
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-08-11 16:18:01
 * @FilePath: \server\system\controller\base\base.go
 */
package base

import (
	"chunDoc/system/core/response"
	"chunDoc/system/service/FileService"

	"github.com/gin-gonic/gin"
)

var this FileService.FileService

//获取文档列表
func List(c *gin.Context) {
	files := this.List()
	response.OkWithData(files, c)
}

/*
 * @Date: 2022-04-12 09:06:05
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-08-11 17:42:15
 * @FilePath: \server\system\controller\base\base.go
 */
package base

import (
	"chunDoc/system/core/request"
	"chunDoc/system/core/response"
	"chunDoc/system/model/RequestModel"
	"chunDoc/system/service/md"

	"github.com/gin-gonic/gin"
)

var this md.MdService

//获取文档列表
func List(c *gin.Context) {
	var r RequestModel.LangName
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	files := this.List(r.LangName)
	response.OkWithData(files, c)
}
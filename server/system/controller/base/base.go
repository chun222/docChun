/*
 * @Date: 2022-04-12 09:06:05
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-08-17 16:47:04
 * @FilePath: \server\system\controller\base\base.go
 */
package base

import (
	"chunDoc/system/core/request"
	"chunDoc/system/core/response"
	"chunDoc/system/model/RequestModel"
	"chunDoc/system/service/md"

	"chunDoc/system/service/configService"

	"github.com/gin-gonic/gin"
)

var thisMd md.MdService

//所有配置
func AllConfig(c *gin.Context) {
	var thisConfig configService.ConfigService
	data := thisConfig.InitConfig()
	response.OkWithData(data, c)
}

//获取文档列表
func List(c *gin.Context) {
	var r RequestModel.InitData
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	files := thisMd.List(r.Lang, r.Version)
	response.OkWithData(files, c)
}

func ReadContent(c *gin.Context) {
	var r RequestModel.Path
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	files := thisMd.ReadContent(r.Path)
	response.OkWithData(files, c)
}

//获取文档列表
func Search(c *gin.Context) {
	var r RequestModel.Search
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	files := thisMd.Search(r.Lang, r.Version, r.Keyword)
	response.OkWithData(files, c)
}

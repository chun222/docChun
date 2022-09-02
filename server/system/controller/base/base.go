/*
 * @Date: 2022-04-12 09:06:05
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-09-02 16:07:24
 * @FilePath: \server\system\controller\base\base.go
 */
package base

import (
	"chunDoc/system/core/request"
	"chunDoc/system/core/response"
	"chunDoc/system/model/RequestModel"
	"chunDoc/system/service/md"
	"chunDoc/system/util/sys"
	"fmt"

	"chunDoc/system/service/configService"

	"github.com/gin-gonic/gin"
)

var thisMd md.MdService
var basedir = sys.ExecutePath() + "/" //根目录
var mdDir = basedir + "md/"

//所有配置
func AllConfig(c *gin.Context) {
	var thisConfig configService.ConfigService
	data := thisConfig.InitConfig()
	response.OkWithData(data, c)
}

//保存配置
func SaveConfigs(c *gin.Context) {
	var thisConfig configService.ConfigService
	var r configService.InitConfigStruct
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	err = thisConfig.SaveConfigs(r)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	response.Ok(c)
}

//获取文档列表
func List(c *gin.Context) {
	var r RequestModel.InitData
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	files := thisMd.List(r)
	response.OkWithData(files, c)
}

func ReadContent(c *gin.Context) {
	var r RequestModel.Path
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}

	files := thisMd.ReadContent(fmt.Sprintf("%s%s/%s/%s/%s", mdDir, r.Project, r.Version, r.Lang, r.Page))
	response.OkWithData(files, c)
}

func SaveContent(c *gin.Context) {
	var r RequestModel.Path
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	err = thisMd.SaveContent(r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func CreateOrUpdateFile(c *gin.Context) {
	var r RequestModel.CreateFileAttr
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	err = thisMd.CreateOrUpdateFileDir(r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

//删除  删除指定目录里的
func RemoveFile(c *gin.Context) {
	var r RequestModel.Remove
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	err = thisMd.RemoveFile(basedir + r.Path)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

//上传文件
func Upload(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	p, err := thisMd.Upload(file, c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(p, c)
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

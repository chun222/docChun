/*
 * @Date: 2022-04-12 09:17:10
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-09-02 16:09:27
 * @FilePath: \server\system\router\base.go
 */

package router

import (
	"chunDoc/system/controller/base"
	//"chunDoc/system/middleware"
	"github.com/gin-gonic/gin"
)

func BaseRouter(r *gin.Engine) {
	//无需登录
	r.POST("/doclist", base.List)
	r.POST("/read", base.ReadContent)
	r.POST("/search", base.Search)
	r.POST("/allconfig", base.AllConfig)
	r.POST("/saveconfigs", base.SaveConfigs)

	r.POST("/savecontent", base.SaveContent)
	r.POST("/create_update_file", base.CreateOrUpdateFile)

	r.POST("/removefile", base.RemoveFile)

	r.POST("/uploadfile", base.Upload)

	// authG := r.Group("tools")
	// //只需要登录
	// authG.Use(middleware.JwtMid())
	// {
	// 	authG.POST("/sequence", tools.Sequence)
	// 	authG.POST("/genpdf", pdf.GenPdf)
	// }

}

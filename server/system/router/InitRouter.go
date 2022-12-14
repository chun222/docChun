/*
 * @Date: 2022-02-14 10:46:28
 * @LastEditors: 春贰
 * @Desc:
 * @LastEditTime: 2022-09-02 17:43:29
 * @FilePath: \server\system\router\InitRouter.go
 */

package router

import (
	"embed"
	"io/fs"
	"net/http"

	"chunDoc/system/core/config"
	"chunDoc/system/middleware"
	"chunDoc/system/util/sys"

	"github.com/gin-gonic/gin"
)

func InitRouter(staticFs embed.FS) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	basedir := sys.ExecutePath() + "\\" //根目录
	r := gin.New()

	r.MaxMultipartMemory = int64(config.Instance().Upload.MaxSize * 1024 * 1024 * 2)

	r.Static("/runtime/file", basedir+"runtime/file")
	//r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/pages"
		r.HandleContext(c)
	})

	r.Static("/static", basedir+"static") //静态文件
	fads, _ := fs.Sub(staticFs, "pages")
	r.StaticFS("/pages", http.FS(fads)) //挂载到二进制中

	r.StaticFS("/apidoc", http.Dir(basedir+"apidoc"))
	//挂载目录提供http访问
	r.Static("/upload", basedir+config.Instance().Upload.Path)

	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())

	//全局中间件
	r.Use(middleware.CorsMid())
	//注册基础路由
	BaseRouter(r)

	return r
}

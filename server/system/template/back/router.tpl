/* 
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:自动生成  
 */
package router

import (
	 "chunDoc/app/controller/{{.TableName}}"
	"chunDoc/system/middleware"
	"github.com/gin-gonic/gin"
)

func R_{{.TableName}}(r *gin.Engine) {
	authG := r.Group("{{.TableName}}")
	authG.Use(middleware.JwtMid(), middleware.PermissionMid())
	{
		authG.POST("/list", {{.TableName}}.List)
		authG.POST("/add",  {{.TableName}}.Add)
		authG.POST("/edit", {{.TableName}}.Edit)
		authG.POST("/remove", {{.TableName}}.Romove) 

	}

} 
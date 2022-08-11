/*
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:路由入口
 */
package app

import (
	AppDbModel "chunDoc/app/model/DbModel"
	"chunDoc/app/router"
	"chunDoc/app/task"
	"chunDoc/system/core/db"
	"github.com/gin-gonic/gin"
)

func InitAppRouter(r *gin.Engine) *gin.Engine { 
	{{range .Routers }}router.R_{{.}}(r) 
	{{ end }}
	return r
}

func InitAppDb() {  
	 db.CheckTableData({{range .Dbtables }}&AppDbModel.{{.}}{},
	 {{ end }})
}

func InitAppTask() {
	task.Init()
}

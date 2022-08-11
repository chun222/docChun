/*
 * @Date: 2022-08-01 14:47:52
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:启动核心服务
 * @LastEditTime: 2022-08-11 16:33:19
 * @FilePath: \server\system\init.go
 */
package system

import (
	"chunDoc/system/core/config"
	"chunDoc/system/core/log"
	"chunDoc/system/core/task"
	"chunDoc/system/core/trans"
	"chunDoc/system/router"
	"chunDoc/system/util/sys"
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	//"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func Init(staticFs embed.FS) {
	basedir := sys.ExecutePath() + "\\" //根目录
	config.InitConfig(basedir + "config.toml")
	log.InitLog()         //初始化日志
	trans.InitTrans("en") //初始化翻译
	r := router.InitRouter(staticFs)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Instance().App.HttpPort),
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	url := fmt.Sprintf(`http://127.0.0.1:%d`, config.Instance().App.HttpPort)
	fmt.Println("运行地址：" + url)
	// exec.Command("cmd", "/C", "start "+url).Run() //windows打开默认浏览器
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Write(log.Error, err.Error())
			os.Exit(0)
		}
	}()
	go task.Init() //初始化system任务服务

	shutDown(s)
}

func shutDown(s *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Write(log.Fatal, "服务关闭:"+err.Error())
	}
	fmt.Println("退出服务")
}

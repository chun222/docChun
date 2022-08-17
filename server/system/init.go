/*
 * @Date: 2022-08-01 14:47:52
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:启动核心服务
 * @LastEditTime: 2022-08-17 10:58:48
 * @FilePath: \server\system\init.go
 */
package system

import (
	"chunDoc/system/common/initial"
	"chunDoc/system/core/config"
	"chunDoc/system/core/log"
	"chunDoc/system/core/task"
	"chunDoc/system/core/trans"
	"chunDoc/system/router"
	"chunDoc/system/util/file"
	"chunDoc/system/util/sys"

	"chunDoc/system/service/configService"
	"context"
	"embed"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	//"os/exec"
	"os/signal"
	"syscall"
	"time"
)

var basedir = sys.ExecutePath() + "/" //根目录

func Init(staticFs embed.FS) {
	InitData()
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

//初始化data，json格式
func InitData() {

	initlockFilePtah := basedir + "init_lock"
	if file.CheckNotExist(initlockFilePtah) {

		file.IsNotExistMkDir(basedir + "data")
		//写入多语言配置
		langPath := basedir + "data/lang.json"
		if file.CheckNotExist(langPath) {
			err := ioutil.WriteFile(langPath, []byte(initial.DataLang), 0777)
			if err != nil {
				log.Write(log.Fatal, "lang:"+err.Error())
			}
		}

		//写入多版本配置
		versionPath := basedir + "data/version.json"
		if file.CheckNotExist(versionPath) {
			err := ioutil.WriteFile(versionPath, []byte(initial.DataVersion), 0777)
			if err != nil {
				log.Write(log.Fatal, "version:"+err.Error())
			}
		}

		//读取默认版本
		var serConfig configService.ConfigService
		var versions []configService.DirAlias
		var langs []configService.DirAlias
		serConfig.GetStructBystring([]byte(initial.DataVersion), &versions)
		serConfig.GetStructBystring([]byte(initial.DataLang), &langs)

		for _, v := range versions {
			for _, l := range langs {
				dirMk := fmt.Sprintf("%s/md/%s/%s", basedir, v.Dir, l.Dir)
				file.IsNotExistMkDir(dirMk)
				fileMk := fmt.Sprintf("%s/index.md", dirMk)
				if file.CheckNotExist(fileMk) {
					contentInit := initial.EnMd
					if l.Dir == "zh-cn" {
						contentInit = initial.ZhMd
					}
					err := ioutil.WriteFile(fileMk, []byte(contentInit), 0777)
					if err != nil {
						log.Write(log.Fatal, fileMk+err.Error())
					}
				}
			}
		}

		if lockfile, err := os.Create(initlockFilePtah); err != nil {
			panic(err)
		} else {
			lockfile.Close()
		}
	}

}

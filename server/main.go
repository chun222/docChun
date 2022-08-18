/*
 * @Date: 2022-02-14 10:20:39
 * @LastEditors: 春贰
 * @Desc:
 * @LastEditTime: 2022-08-18 14:30:14
 * @FilePath: \server\main.go
 */
//go:generate goversioninfo
package main

import (
	"chunDoc/system"
	"embed"
	"fmt"
	"github.com/kardianos/service"
	"log"
	"os"
)

//go:embed pages
var pagesFs embed.FS

var serviceConfig = &service.Config{
	Name:        "chunDocService",
	DisplayName: "chunDocService",
	Description: "markdown在线生成文件管理服务，主页：https://github.com/chun222",
}

func main() {

	// 构建服务对象
	prog := &Program{}

	//注册到服务的时候，需要设置目录
	if len(os.Args) >= 2 {
		typecmd := os.Args[1]
		if typecmd == "install" {
			dir, _ := os.Getwd()
			serviceConfig.Arguments = append(serviceConfig.Arguments, dir)
		}
	}

	s, err := service.New(prog, serviceConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 用于记录系统日志
	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
		return
	}

	cmd := os.Args[1]

	if cmd == "install" {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("安装成功")
	} else if cmd == "uninstall" {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("卸载成功")
	} else if cmd == "start" {
		err = s.Start()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("开始成功")
	} else if cmd == "stop" {
		err = s.Stop()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("停止成功")
	} else {
		//直接运行 第一个参数为运行目录
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
		return
	}

	// install, uninstall, start, stop 的另一种实现方式,这里不适用
	// err = service.Control(s, os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println(os.Args[1], "成功")
	// }
}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	log.Println("开始服务")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	log.Println("停止服务")
	return nil
}

func (p *Program) run() {
	log.Println("开始成功")
	system.Init(pagesFs)
}

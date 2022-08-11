/*
 * @Date: 2022-02-14 10:20:39
 * @LastEditors: 春贰
 * @Desc:
 * @LastEditTime: 2022-08-11 16:27:22
 * @FilePath: \server\main.go
 */
package main

import (
	"chunDoc/system"
	"embed"
	"fmt"
	"github.com/kardianos/service"
	"log"
	"os"
)

//go:embed static
var staticFs embed.FS

var serviceConfig = &service.Config{
	Name:        "TjsEmsServer_V2_",
	DisplayName: "TjsEmsServer_V2_",
	Description: "天俱时能源管理服务2.0",
}

func main() {

	// 构建服务对象
	prog := &Program{}

	//注册到服务的时候，需要设置目录
	if len(os.Args) >= 3 {
		serviceConfig.Arguments = append(serviceConfig.Arguments, os.Args[2])
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
		fmt.Println("卸载成功")
	} else if cmd == "stopp" {
		err = s.Start()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("卸载成功")
	} else {
		//直接运行 第一个参数为运行目录
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
		return
	}

	// install, uninstall, start, stop 的另一种实现方式
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
	system.Init(staticFs)
}

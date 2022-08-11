/*
 * @Date: 2022-05-07 11:18:11
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-06-09 17:19:12
 * @FilePath: \server\system\core\task\task.go
 */
package task

import (
	"fmt"
	"runtime"
	"time"

	"chunDoc/system/core/log"
)

func Init() {
	//避免影响主程序
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("task发生错误", err)
		}
	}()

	go WorkOrderTask()

}

//工单任务
func WorkOrderTask() {
	//避免影响主程序
	defer func() {
		if err := recover(); err != nil {
			log.Write(log.Error, "WorkOrderTask发生错误")
			fmt.Println("WorkOrderTask发生错误", err)
		}
	}() //捕获错误
	for {
		//执行任务
		time.Sleep(5 * time.Minute)
		fmt.Println("goroutine数量：", runtime.NumGoroutine())
	}
}

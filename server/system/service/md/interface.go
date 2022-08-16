/*
 * @Date: 2022-08-15 10:23:54
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-08-16 09:08:33
 * @FilePath: \server\system\service\md\interface.go
 */
package md

type IDoc interface {
	ReadContent(path string)            //读取内容接口
	Search(lang string, keyword string) //搜索
}

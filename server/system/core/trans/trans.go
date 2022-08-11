/*
 * @Date: 2022-02-25 08:41:53
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 翻译核心
 * @LastEditTime: 2022-02-25 10:35:05
 * @FilePath: \gin-antd\server\app\core\trans\trans.go
 */
package trans

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'

		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		Trans, ok = uni.GetTranslator(locale)

		if !ok {
			fmt.Println("uni.GetTranslator failed", locale)
			os.Exit(0)
		}

		// 注册翻译器
		switch locale {
		case "en":
			enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		return
	} else {

		fmt.Println("uni.GetTranslator failed", locale)
		os.Exit(0)

	}
	return
}

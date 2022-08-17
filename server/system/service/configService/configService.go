/*
 * @Date: 2022-08-17 09:52:04
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-08-17 10:59:32
 * @FilePath: \server\system\service\configService\configService.go
 */
package configService

import (
	"chunDoc/system/util/sys"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ConfigService struct {
}

type DirAlias struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
}

var basedir = sys.ExecutePath() + "/" //根目录

var LangDataPath = fmt.Sprintf("%sdata/lang.json", basedir)
var VersionDataPath = fmt.Sprintf("%sdata/version.json", basedir)

//读取文件夹别名配置
func (_this *ConfigService) GetStructBystring(s []byte, r *[]DirAlias) error {
	return json.Unmarshal(s, r)
}

//初始化配置
type InitConfigStruct struct {
	Lang    []DirAlias
	Version []DirAlias
}

func (_this *ConfigService) InitConfig() InitConfigStruct {
	var re InitConfigStruct
	langBytes, err := ioutil.ReadFile(LangDataPath)
	if err == nil {
		_this.GetStructBystring(langBytes, &re.Lang)
	}

	versionBytes, err := ioutil.ReadFile(VersionDataPath)
	if err == nil {
		_this.GetStructBystring(versionBytes, &re.Version)
	}

	return re
}

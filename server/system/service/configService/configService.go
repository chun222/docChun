/*
 * @Date: 2022-08-17 09:52:04
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-08-26 10:58:54
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
	DataFile []DataFile
}

type DataFile struct {
	Name     string `json:"name"`
	FullPath string `json:"fullpath"`
}

type DirAlias struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
}

var basedir = sys.ExecutePath() + "/" //根目录

//返回获取的所有配置路径
func (_this *ConfigService) AllConfigPaths() {
	ProjectDataPath := fmt.Sprintf("%sdata/project.json", basedir)
	_this.DataFile = append(_this.DataFile, DataFile{
		Name:     "project",
		FullPath: ProjectDataPath,
	})
	LangDataPath := fmt.Sprintf("%sdata/lang.json", basedir)
	_this.DataFile = append(_this.DataFile, DataFile{
		Name:     "lang",
		FullPath: LangDataPath,
	})
	VersionDataPath := fmt.Sprintf("%sdata/version.json", basedir)
	_this.DataFile = append(_this.DataFile, DataFile{
		Name:     "version",
		FullPath: VersionDataPath,
	})
}

//读取文件夹别名配置
func (_this *ConfigService) GetStructBystring(s []byte, r *[]DirAlias) error {
	return json.Unmarshal(s, r)
}

//初始化配置
type InitConfigStruct struct {
	Lang    []DirAlias
	Version []DirAlias
	Project []DirAlias
}

func (_this *ConfigService) InitConfig() InitConfigStruct {

	_this.AllConfigPaths()

	var re InitConfigStruct
	for _, v := range _this.DataFile {

		switch v.Name {
		case "lang":
			versionBytes, err := ioutil.ReadFile(v.FullPath)
			if err == nil {
				_this.GetStructBystring(versionBytes, &re.Lang)
			}
		case "version":
			versionBytes, err := ioutil.ReadFile(v.FullPath)
			if err == nil {
				_this.GetStructBystring(versionBytes, &re.Version)
			}
		case "project":
			versionBytes, err := ioutil.ReadFile(v.FullPath)
			if err == nil {
				_this.GetStructBystring(versionBytes, &re.Project)
			}
		}

	}

	return re
}

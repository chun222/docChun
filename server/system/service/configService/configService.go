/*
 * @Date: 2022-08-17 09:52:04
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-09-01 14:07:57
 * @FilePath: \server\system\service\configService\configService.go
 */
package configService

import (
	"chunDoc/system/util/sys"

	"chunDoc/system/core/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	Langs    []DirAlias
	Versions []DirAlias
	Projects []DirAlias
}

func (_this *ConfigService) InitConfig() InitConfigStruct {

	_this.AllConfigPaths()

	var re InitConfigStruct
	for _, v := range _this.DataFile {

		switch v.Name {
		case "lang":
			versionBytes, err := ioutil.ReadFile(v.FullPath)
			if err == nil {
				_this.GetStructBystring(versionBytes, &re.Langs)
			}
		case "version":
			versionBytes, err := ioutil.ReadFile(v.FullPath)
			if err == nil {
				_this.GetStructBystring(versionBytes, &re.Versions)
			}
		case "project":
			versionBytes, err := ioutil.ReadFile(v.FullPath)
			if err == nil {
				_this.GetStructBystring(versionBytes, &re.Projects)
			}
		}

	}

	return re
}

//保存配置
func (_this *ConfigService) SaveConfigs(configs InitConfigStruct) error {
	_this.AllConfigPaths()
	//遍历配置
	for _, v := range _this.DataFile {
		switch v.Name {
		case "lang":
			if err := _this.writeJsonConfig(configs.Langs, v.FullPath); err != nil {
				log.Write(log.Error, err.Error())
			}
		case "version":
			if err := _this.writeJsonConfig(configs.Versions, v.FullPath); err != nil {
				log.Write(log.Error, err.Error())
			}
		case "project":
			if err := _this.writeJsonConfig(configs.Projects, v.FullPath); err != nil {
				log.Write(log.Error, err.Error())
			}
		}
	}
	return nil
}

//将结构体转json并写入配置
func (_this *ConfigService) writeJsonConfig(in interface{}, path string) error {

	content, err := json.Marshal(in)
	if err != nil {
		return err
	}

	//写入内容
	filewrite, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	filewrite.Write(content)
	// n, _ := filewrite.Seek(0, os.SEEK_END)
	// _, err = filewrite.WriteAt([]byte(headInfo+r.Content), n)
	defer filewrite.Close()
	return nil
}

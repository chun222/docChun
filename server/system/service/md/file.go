/*
 * @Date: 2022-03-02 10:32:10
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:markdown
 * @LastEditTime: 2022-08-11 17:36:07
 * @FilePath: \server\system\service\md\file.go
 */

package md

import (
	"chunDoc/system/util/file"
	"chunDoc/system/util/str"
	"chunDoc/system/util/sys"
	"io/ioutil"
	"sort"
)

type MdService struct{}

var basedir = sys.ExecutePath() + "/" //根目录
var mdDir = basedir + "md/"

func (this *MdService) List(lang string) []SysFile {
	filesr := GetDirFiles(mdDir + lang)
	return filesr
}

type SysFile struct {
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Size       float64   `json:"size"`
	CreateTime string    `json:"create_time"`
	Fullpath   string    `json:"fullpath"`
	Child      []SysFile `json:"children"`
}

//递归获取文件夹下所有内容
func GetDirFiles(path string) []SysFile {
	f := SysFile{
		Name: path,
		Type: "dir",
	}
	r := getFileTree(path, &f)
	return r.Child
}

//递归获取文件夹下所有内容
func getFileTree(path string, s *SysFile) *SysFile {

	fs, _ := ioutil.ReadDir(path)
	sort.Sort(file.ByModTime(fs))
	s.Child = make([]SysFile, 0)
	for i, f := range fs {
		if f.IsDir() {
			s.Child = append(s.Child, SysFile{
				Name:       f.Name(),
				Size:       0,
				Fullpath:   path + f.Name(),
				Type:       "dir",
				CreateTime: f.ModTime().Format("2006-01-02 15:04:05"),
			})
			getFileTree(path+f.Name()+"/", &s.Child[i])

		} else {

			s.Child = append(s.Child, SysFile{
				Name:       str.BeforeLast(f.Name(), "."),
				Size:       float64(f.Size()) / 1024, //返回kb
				Fullpath:   path + f.Name(),
				Type:       str.AfterLast(f.Name(), "."),
				CreateTime: f.ModTime().Format("2006-01-02 15:04:05"),
			})
		}
	}

	return s
}

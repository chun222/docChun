/*
 * @Date: 2022-03-02 10:32:10
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:markdown
 * @LastEditTime: 2022-08-12 11:33:57
 * @FilePath: \server\system\service\md\md.go
 */

package md

import (
	//"chunDoc/system/util/file"
	"bufio"
	"chunDoc/system/util/convert"
	"chunDoc/system/util/str"
	"chunDoc/system/util/sys"
	"io/ioutil"
	"os"

	"github.com/dlclark/regexp2"
	//"sort"
)

type MdService struct{}

type LineConfig struct {
	Position *int64
	Label    *string
}

type SysFile struct {
	Key        uint      `json:"key"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Size       float64   `json:"size"`
	CreateTime string    `json:"create_time"`
	Fullpath   string    `json:"fullpath"`
	Child      []SysFile `json:"children"`
}

var basedir = sys.ExecutePath() + "/" //根目录
var mdDir = basedir + "md/"

func (_this *MdService) List(lang string) []SysFile {
	filesr := _this.GetDirFiles(mdDir + lang)
	return filesr
}

//递归获取文件夹下所有内容
func (_this *MdService) GetDirFiles(path string) []SysFile {
	f := SysFile{
		Name: path,
		Type: "dir",
	}
	_this.getFileTree(path, &f)
	return f.Child
}

//递归获取文件夹下所有内容
func (_this *MdService) getFileTree(path string, s *SysFile) {
	fs, _ := ioutil.ReadDir(path)
	///sort.Sort(file.ByModTime(fs))
	//s.Child = make([]SysFile, 0)
	for i, f := range fs {
		if f.IsDir() {
			s.Child = append(s.Child, SysFile{
				Name:       f.Name(),
				Size:       0,
				Fullpath:   path + f.Name(),
				Type:       "dir",
				CreateTime: f.ModTime().Format("2006-01-02 15:04:05"),
			})
			_this.getFileTree(path+"/"+f.Name()+"/", &s.Child[i])

		} else {
			fileone, err := os.Open(path + "/" + f.Name())
			if err != nil {
				panic(err)
			}
			defer fileone.Close()
			rd := bufio.NewReader(fileone)
			line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

			lineRe := _this.readConfigLine(line)

			name := str.BeforeLast(f.Name(), ".") //默认为文件夹名称
			if lineRe.Label != nil {
				name = *lineRe.Label
			}

			s.Child = append(s.Child, SysFile{
				Name:       name,
				Size:       float64(f.Size()) / 1024, //返回kb
				Fullpath:   path + f.Name(),
				Type:       str.AfterLast(f.Name(), "."),
				CreateTime: f.ModTime().Format("2006-01-02 15:04:05"),
			})
		}
	}
}

//解析一行的内容
func (_this *MdService) readConfigLine(line string) (re LineConfig) {

	reg := regexp2.MustCompile(`label|position`, 0)
	if ok, _ := reg.MatchString(line); ok {
		if data, err := reg.FindStringMatch(line); err == nil {
			if data.String() == "label" {
				reg2 := regexp2.MustCompile(`(?<=^\s*label:).*`, 0)
				//找到label之后内容
				if ok, _ := reg2.MatchString(line); ok {
					if data2, err := reg2.FindStringMatch(line); err == nil {
						label := data2.String()
						re.Label = &label
					}
				}
			} else {
				reg2 := regexp2.MustCompile(`(?<=^\s*position:)\d+`, 0)
				//找到label之后内容
				if ok, _ := reg2.MatchString(line); ok {
					if data2, err := reg2.FindStringMatch(line); err == nil {
						if v, err := convert.Int(data2.String()); err == nil {
							re.Position = &v
						}
					}
				}

			}
		}
	}
	return re
}

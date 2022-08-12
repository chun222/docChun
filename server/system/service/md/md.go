/*
 * @Date: 2022-03-02 10:32:10
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:markdown
 * @LastEditTime: 2022-08-12 17:20:10
 * @FilePath: \server\system\service\md\md.go
 */

package md

import (
	//"chunDoc/system/util/file"
	"bufio"
	"chunDoc/system/util/convert"
	"chunDoc/system/util/file"
	"chunDoc/system/util/str"
	"chunDoc/system/util/sys"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/dlclark/regexp2"
	//"sort"
)

type MdService struct{}

type LineConfig struct {
	Position int64
	Label    string
}

type SysFile struct {
	Key        uint       `json:"key"`
	Position   int64      `json:"position"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Size       float64    `json:"size"`
	CreateTime string     `json:"create_time"`
	Fullpath   string     `json:"fullpath"`
	Child      []*SysFile `json:"children"`
}

var basedir = sys.ExecutePath() + "/" //根目录
var mdDir = basedir + "md/"
var FileKey uint = 0

func (_this *MdService) List(lang string) []*SysFile {
	filesr := _this.GetDirFiles(mdDir + lang + "/")
	return filesr
}

//递归获取文件夹下所有内容
func (_this *MdService) GetDirFiles(path string) []*SysFile {
	FileKey = 0
	f := SysFile{
		Key:  FileKey,
		Name: path,
		Type: "dir",
	}
	_this.getFileTree(path, &f)
	sortFileTree(f.Child)

	return f.Child
}

//递归排序
func sortFileTree(f []*SysFile) {
	sort.SliceStable(f, func(i int, j int) bool {
		return f[i].Position < f[j].Position
	})
	for _, v := range f {
		sortFileTree(v.Child)
	}
}

//递归获取文件夹下所有内容
func (_this *MdService) getFileTree(path string, s *SysFile) {
	fs, _ := ioutil.ReadDir(path)
	//sort.Sort(file.ByModTime(fs))
	s.Child = make([]*SysFile, 0)
	var posintion int64 = 999
	for _, f := range fs {
		FileKey++
		if f.IsDir() {
			name := f.Name() //默认为文件夹名称
			configpath := path + "/" + f.Name() + "/_.config"
			if !file.CheckNotExist(configpath) {
				lineRe := _this.readFileHead(configpath)

				fmt.Println("===", configpath, lineRe.Label, lineRe.Position)
				if lineRe.Label != "" {
					name = lineRe.Label
				}
				if lineRe.Position != 0 {
					posintion = lineRe.Position
				}
			}
			fileNow := SysFile{
				Key:        FileKey,
				Position:   posintion,
				Name:       name,
				Size:       0,
				Fullpath:   path + f.Name(),
				Type:       "dir",
				CreateTime: f.ModTime().Format("2006-01-02 15:04:05"),
			}
			s.Child = append(s.Child, &fileNow)
			_this.getFileTree(path+"/"+f.Name()+"/", &fileNow)

		} else {

			fileType := str.AfterLast(f.Name(), ".")
			//只取markdown文件
			if fileType == "md" {

				lineRe := _this.readFileHead(path + "/" + f.Name())
				name := str.BeforeLast(f.Name(), ".") //默认为文件名称
				if lineRe.Label != "" {
					name = lineRe.Label
				}
				if lineRe.Position != 0 {
					posintion = lineRe.Position
				}
				fileNow := SysFile{
					Key:        FileKey,
					Position:   posintion,
					Name:       name,
					Size:       float64(f.Size()) / 1024, //返回kb
					Fullpath:   path + f.Name(),
					Type:       fileType,
					CreateTime: f.ModTime().Format("2006-01-02 15:04:05"),
				}
				s.Child = append(s.Child, &fileNow)
			}
		}
	}
}

//切片的前序元素添加（头部添加）
func preInsertSlice(s []*SysFile, i *SysFile) []*SysFile {
	res := append([]*SysFile{i}, s...)
	return res
}

//解析头部两行的信息
func (_this *MdService) readFileHead(fullpath string) LineConfig {
	var re LineConfig = LineConfig{
		Position: 999,
		Label:    "",
	}
	fileone, err := os.Open(fullpath)
	if err != nil {
		return re
	}
	defer fileone.Close()
	rd := bufio.NewReader(fileone)

	//读取2行
	for i := 0; i < 2; i++ {

		//if line, err := rd.ReadString('\n'); err == nil {
		line, _ := rd.ReadString('\n') //这里报错可以忽略
		labelReturn, positionReturn := _this.readConfigLine(line)
		if positionReturn != nil {
			re.Position = *positionReturn
		}
		if labelReturn != nil {
			re.Label = *labelReturn
		}

	}
	return re
}

//解析一行的内容
func (_this *MdService) readConfigLine(line string) (*string, *int64) {
	var labelReturn *string
	var positionReturn *int64

	reg := regexp2.MustCompile(`label|position`, 0)
	if ok, _ := reg.MatchString(line); ok {
		if data, err := reg.FindStringMatch(line); err == nil {
			if data.String() == "label" {
				reg2 := regexp2.MustCompile(`(?<=^\s*label:).*`, 0)
				//找到label之后内容
				if ok, _ := reg2.MatchString(line); ok {
					if data2, err := reg2.FindStringMatch(line); err == nil {
						label := data2.String()
						labelReturn = &label
					}
				}
			} else {
				reg2 := regexp2.MustCompile(`(?<=^\s*position:)\d+`, 0)
				//找到position之后内容
				if ok, _ := reg2.MatchString(line); ok {
					if data2, err := reg2.FindStringMatch(line); err == nil {
						if v, err := convert.Int(data2.String()); err == nil {
							positionReturn = &v
						}
					}
				}

			}
		}
	}
	return labelReturn, positionReturn
}

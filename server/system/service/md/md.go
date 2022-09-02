/*
 * @Date: 2022-03-02 10:32:10
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:markdown
 * @LastEditTime: 2022-09-02 16:33:55
 * @FilePath: \server\system\service\md\md.go
 */

package md

import (
	//"chunDoc/system/util/file"
	"bufio"
	"chunDoc/system/core/config"
	"chunDoc/system/core/log"
	"chunDoc/system/model/RequestModel"
	"chunDoc/system/util/convert"
	"chunDoc/system/util/file"
	"chunDoc/system/util/str"
	"chunDoc/system/util/sys"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	//"sort"
)

type MdService struct {
}

type LineConfig struct {
	Position   int64
	Label      string
	EffectRows int //有效行数
}

type SysFile struct {
	Key          uint          `json:"key"`
	Position     int64         `json:"position"`
	Realname     string        `json:"realname"`
	Name         string        `json:"name"`
	Type         FileTypeAllow `json:"type"`
	Size         float64       `json:"size"`
	CreateTime   string        `json:"create_time"`
	Fullpath     string        `json:"fullpath"`
	Relativepath string        `json:"relativepath"`
	Child        []*SysFile    `json:"children"`
}

type FileTypeAllow string

const TypeDir FileTypeAllow = "dir"
const TypeMd FileTypeAllow = "md"

var basedir = sys.ExecutePath() + "/" //根目录
var mdDir = "md/"
var FileKey uint = 0

var configFileName = "_.config"

func (_this *MdService) List(r RequestModel.InitData) []*SysFile {
	filesr := _this.getDirFiles(fmt.Sprintf("%s%s/%s/%s/", mdDir, r.Project, r.Version, r.Lang))
	return filesr
}

//递归获取文件夹下所有内容
func (_this *MdService) getDirFiles(path string) []*SysFile {
	FileKey = 0
	f := SysFile{
		Key:  FileKey,
		Name: path,
		Type: TypeDir,
	}
	_this.getFileTree(path, &f, path)
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

//递归获取文件夹下所有内容 , path 不包含根目录 , rePath:相对根目录
func (_this *MdService) getFileTree(path string, s *SysFile, rePath string) {
	fs, _ := ioutil.ReadDir(basedir + path)
	//sort.Sort(file.ByModTime(fs))
	s.Child = make([]*SysFile, 0)
	var posintion int64 = 999
	for _, f := range fs {
		FileKey++
		if f.IsDir() {
			name := f.Name() //默认为文件夹名称
			configpath := path + "/" + f.Name() + "/" + configFileName
			if !file.CheckNotExist(configpath) {
				lineRe := _this.readFileHead(configpath)
				if lineRe.Label != "" {
					name = lineRe.Label
				}
				if lineRe.Position != 0 {
					posintion = lineRe.Position
				}
			}
			fileNow := SysFile{
				Key:          FileKey,
				Position:     posintion,
				Name:         name,
				Realname:     f.Name(),
				Size:         0,
				Fullpath:     path + f.Name(),
				Relativepath: f.Name(),
				Type:         "dir",
				CreateTime:   f.ModTime().Format("2006-01-02 15:04:05"),
			}
			s.Child = append(s.Child, &fileNow)
			_this.getFileTree(path+"/"+f.Name()+"/", &fileNow, rePath)

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
					Key:          FileKey,
					Position:     posintion,
					Name:         name,
					Realname:     str.BeforeLast(f.Name(), "."),
					Size:         float64(f.Size()) / 1024, //返回kb
					Fullpath:     path + f.Name(),
					Relativepath: str.After(path, rePath) + f.Name(),
					Type:         TypeMd,
					CreateTime:   f.ModTime().Format("2006-01-02 15:04:05"),
				}
				s.Child = append(s.Child, &fileNow)
			}
		}
	}
}

//获取文件列表list，不包含文件夹
func (_this *MdService) getFileList(path string) []SysFile {

	fs, _ := ioutil.ReadDir(basedir + path)
	//sort.Sort(file.ByModTime(fs))
	result := make([]SysFile, 0)
	var posintion int64 = 999
	for _, f := range fs {
		if f.IsDir() {
			result = append(result, _this.getFileList(path+"/"+f.Name()+"/")...)

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
					Type:       TypeMd,
					CreateTime: f.ModTime().Format("2006-01-02 15:04:05"),
				}
				result = append(result, fileNow)
			}
		}
	}

	return result
}

//切片的前序元素添加（头部添加）
func preInsertSlice(s []*SysFile, i *SysFile) []*SysFile {
	res := append([]*SysFile{i}, s...)
	return res
}

//解析头部两行的信息
func (_this *MdService) readFileHead(fullpath string) LineConfig {
	var re LineConfig = LineConfig{
		Position:   999,
		Label:      "",
		EffectRows: 0,
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
			re.EffectRows++
		}
		if labelReturn != nil {
			re.Label = *labelReturn
			re.EffectRows++
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

//保存md内容
func (_this *MdService) SaveContent(r RequestModel.Path) error {

	fullpath := fmt.Sprintf("%s%s%s/%s/%s/%s", basedir, mdDir, r.Project, r.Version, r.Lang, r.Page)
	fileone, err := os.Open(fullpath)
	if err != nil {
		return err
	}
	defer fileone.Close()
	buf := bufio.NewReader(fileone)

	//读取2行,判断是否有位置和排序信息
	headInfo := ""
	for i := 0; i < 2; i++ {
		line, _ := buf.ReadString('\n')
		s, in := _this.readConfigLine(line)
		if s != nil || in != nil {
			headInfo = headInfo + line
		}
	}
	//写入内容
	filewrite, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	n, _ := filewrite.Seek(0, os.SEEK_END)
	_, err = filewrite.WriteAt([]byte(headInfo+r.Content), n)
	defer filewrite.Close()
	return err

}

//读取md内容
func (_this *MdService) ReadContent(fullpath string) string {
	//fullpath := fmt.Sprintf("%s%s/%s/%s/%s", mdDir, r.Project, r.Version, r.Lang, r.Page)
	re := ""
	fileone, err := os.Open(fullpath)
	if err != nil {
		return re
	}
	defer fileone.Close()
	buf := bufio.NewReader(fileone)

	//读取2行,判断是否返回前两行
	i := 0
	for {
		line, err := buf.ReadString('\n')
		if i < 2 {
			if s, in := _this.readConfigLine(line); s == nil && in == nil {
				re = re + line
			}
		} else {
			re = re + line
		}
		i++

		if err != nil {
			if err == io.EOF { //读取结束，会报EOF
				break
			}
			break
		}
	}
	return re

}

//创建文件或文件夹并赋予信息
func (_this *MdService) createFileDir(r RequestModel.CreateFileAttr) error {
	//文件不存在
	fullpath := basedir + r.Parent + "/" + r.Dir
	headinfo := fmt.Sprintf("label:%s\nposition:%d\n", r.Name, r.Position)
	if file.CheckNotExist(fullpath) {
		if r.Type == string(TypeMd) {
			//创建文件并写入
			fullpath = fullpath + ".md" //加后缀
			err := ioutil.WriteFile(fullpath, []byte(headinfo), 0777)
			if err != nil {

				log.Write(log.Error, err.Error())
				return err
			}
			return nil
		} else {
			//文件夹
			err := file.MkDir(fullpath)
			if err != nil {
				log.Write(log.Error, err.Error())
				return err
			}
			//创建文件信息
			configPath := fullpath + "/" + configFileName
			err = ioutil.WriteFile(configPath, []byte(headinfo), 0777)
			if err != nil {
				log.Write(log.Error, err.Error())
				return err
			}
			return nil
		}
	} else {
		return fmt.Errorf("文件已存在")
	}
}

func (_this *MdService) updateFileDir(r RequestModel.CreateFileAttr) error {
	fullpath := basedir + r.Path
	headinfo := fmt.Sprintf("label:%s\nposition:%d\n", r.Name, r.Position)

	//文件存在
	if !file.CheckNotExist(fullpath) {
		//md修改前两行数据
		if r.Type == string(TypeMd) {

			mdContent := _this.ReadContent(fullpath)
			//覆盖写入
			filewrite, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
			if err != nil {
				log.Write(log.Error, err.Error())
				return err
			}
			_, err = filewrite.WriteString(headinfo + mdContent)
			if err != nil {
				log.Write(log.Error, err.Error())
				return err
			}
			defer filewrite.Close()
		} else {

			configPath := fullpath + "/" + configFileName
			//覆盖写入
			filewrite, err := os.OpenFile(configPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
			if err != nil {
				log.Write(log.Error, err.Error())
				return err
			}
			_, err = filewrite.WriteString(headinfo)
			if err != nil {
				log.Write(log.Error, err.Error())
				return err
			}
			defer filewrite.Close()
		}

		return nil
	} else {
		return fmt.Errorf("文件不存在了,修改失败")
	}
}

//创建或更新文件属性
func (_this *MdService) CreateOrUpdateFileDir(r RequestModel.CreateFileAttr) error {

	//根据path是否传入判断是创建还是更新
	if r.Path == "" {
		return _this.createFileDir(r)
	} else {
		//更新
		return _this.updateFileDir(r)
	}
}

//删除文件
func (_this *MdService) RemoveFile(path string) error {
	return os.RemoveAll(path)
}

//上传文件
func (_this *MdService) Upload(r *multipart.FileHeader, c *gin.Context) (string, error) {
	// 上传文件到指定的路径
	if r.Size > int64(config.Instance().Upload.MaxSize*1024*1024) {
		return "", fmt.Errorf("file too large")
	}
	dateDir := time.Now().Format("20060102")
	uploaddir := fmt.Sprintf("%s%s/%s/", basedir, config.Instance().Upload.Path, dateDir)
	file.MkDir(uploaddir)
	uid := uuid.New().String()
	filename := uid + "." + str.AfterLast(r.Filename, ".")
	filepath := uploaddir + filename
	if file.CheckNotExist(filepath) {
		err := c.SaveUploadedFile(r, filepath)
		if err != nil {
			return "", err
		} else {
			//返回路径
			returnPath := fmt.Sprintf("upload/%s/%s", dateDir, filename)
			return returnPath, nil
		}
	} else {
		return "", fmt.Errorf("has exist file")
	}

}

//搜索

type FindedResult struct {
	Id       string
	PageName string
	PagePath string
	Text     string
}

func (_this *MdService) Search(lang string, version string, keyword string) []FindedResult {
	files := _this.getFileList(fmt.Sprintf("%s%s%s/%s/", basedir, mdDir, version, lang))
	var result []FindedResult
	for _, vfile := range files {
		f, err := os.Open(basedir + vfile.Fullpath)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		rd := bufio.NewReader(f)

		HasIndexTitles := make(map[string]int, 0)
		IndexTitle := ""
		HasTitleFind := false

		for {
			line, err := rd.ReadString('\n')                              //以'\n'为结束符读入一行
			reg := regexp2.MustCompile(`(?<=^\s*#+\s+)[^\s]+.*[^\s$]`, 0) //匹配标题
			if ok, _ := reg.MatchString(line); ok {
				if data, err := reg.FindStringMatch(line); err == nil {
					regReplace := regexp2.MustCompile(`[\*\[\]]`, 0) //替换的
					IndexTitle, _ = regReplace.Replace(data.String(), "", -1, -1)
					IndexTitle = strings.ReplaceAll(IndexTitle, " ", "-")

					//判断标题是否找到，如果找到 下面的内容不继续搜索
					if find := strings.Contains(IndexTitle, keyword); find {
						HasTitleFind = true
					} else {
						HasTitleFind = false
					}

					if num, ok := HasIndexTitles[IndexTitle]; ok {
						HasIndexTitles[IndexTitle] = num + 1
						IndexTitle = fmt.Sprintf("%s-%d", IndexTitle, num+1)
					} else {
						HasIndexTitles[IndexTitle] = 1
					}

					if HasTitleFind {
						result = append(result, FindedResult{
							Id:       IndexTitle,
							PageName: vfile.Name,
							PagePath: vfile.Fullpath,
							Text:     line,
						})
					}

				}
			} else {
				//非标题
				if find := strings.Contains(line, keyword); find && !HasTitleFind {
					result = append(result, FindedResult{
						Id:       IndexTitle,
						PageName: vfile.Name,
						PagePath: vfile.Fullpath,
						Text:     line,
					})
				}
			}

			if err != nil {
				// if err == io.EOF { //读取结束，会报EOF
				// 	break
				// }
				break
			}
		}

	}

	return result
}

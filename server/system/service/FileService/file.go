/*
 * @Date: 2022-03-02 10:32:10
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:备份
 * @LastEditTime: 2022-08-11 16:23:50
 * @FilePath: \server\system\service\FileService\file.go
 */

package FileService

import (
	"errors"
	"os"
	"strings"

	"chunDoc/system/core/config"
	"chunDoc/system/model/RequestModel"
	"chunDoc/system/util/file"
	"chunDoc/system/util/sys"

	"github.com/gin-gonic/gin"
)

type FileService struct{}

var uploadpath = sys.ExecutePath() + "\\" + strings.TrimRight(config.Instance().Upload.Path, "/") + "/"

func (fileservice *FileService) List() []file.SysFile {
	filesr := file.GetDirFiles(uploadpath)
	for i, v := range filesr {
		name := v.Name + "." + v.Type
		if v.Type == "dir" {
			name = v.Name
		}
		filesr[i].Fullpath = strings.TrimRight(config.Instance().Upload.Path, "/") + "/" + name
	}
	return filesr
}

func (fileservice *FileService) Upload(r RequestModel.FileUpload, c *gin.Context) error {
	// 上传文件到指定的路径
	if r.File.Size > int64(config.Instance().Upload.MaxSize*1024*1024) {
		return errors.New("file too large")
	}

	uploaddir := uploadpath
	for i, v := range r.Path {
		if i > 0 {
			uploaddir = uploaddir + v + "/"
		}
	}
	file.MkDir(uploaddir)
	if file.CheckNotExist(uploaddir + r.File.Filename) {
		return c.SaveUploadedFile(&r.File, uploaddir+r.File.Filename)
	} else {
		return errors.New("has exist file")
	}

}

//创建文件夹
func (fileservice *FileService) CreateDir(r RequestModel.FileUpload) {
	// 上传文件到指定的路径
	uploaddir := uploadpath
	for i, v := range r.Path {
		if i > 0 {
			uploaddir = uploaddir + v + "/"
		}
	}
	file.MkDir(uploaddir + r.Dirname)
}

func (fileservice *FileService) Remove(r RequestModel.FileUpload) error {
	return os.RemoveAll(r.Dirname)
}

/*
 * @Date: 2022-02-14 14:20:21
 * @LastEditors: 春贰
 * @Desc:
 * @LastEditTime: 2022-03-31 10:58:40
 * @FilePath: \server\system\util\file\file.go
 */
package file

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"sort"

	"chunDoc/system/util/str"
)

//时间排序
type ByModTime []fs.FileInfo

func (fis ByModTime) Len() int {
	return len(fis)
}
func (fis ByModTime) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}
func (fis ByModTime) Less(i, j int) bool {
	return fis[i].ModTime().After(fis[j].ModTime())
}

//获取文件夹所有文件
func GetDirfileNames(dir string) []string {
	var filenames []string
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		filenames = append(filenames, str.Before(f.Name(), "."))
	}
	return filenames

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

	if !CheckNotExist(path) {
		fs, _ := ioutil.ReadDir(path)
		sort.Sort(ByModTime(fs))
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
	}
	return s
}

// GetSize get the file size
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetExt get the file ext
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckNotExist check if the file exists
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// CheckPermission check if the file has permission
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir create a directory if it does not exist
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

// MkDir create a directory
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open a file according to a specific mode
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// MustOpen maximize trying to open the file
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}

// IsDir determines whether the specified path is a directory.
func IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}
	if nil != err {
		fmt.Printf("Determines whether [%s] is a directory failed: [%v]", path, err)
		return false
	}
	return fio.IsDir()
}

// CopyFile copies the source file to the dest file.
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourcefile.Close()
	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destfile.Close()
	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}
	}
	return nil
}

// CopyDir copies the source directory to the dest directory.
func CopyDir(source string, dest string) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}
	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}
	directory, err := os.Open(source)
	if err != nil {
		return err
	}
	defer directory.Close()
	objects, err := directory.Readdir(-1)
	if err != nil {
		return err
	}
	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(srcFilePath, destFilePath)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(srcFilePath, destFilePath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

/*
 * @Date: 2022-06-01 16:15:51
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-07-15 16:28:26
 * @FilePath: \server\system\util\pdf\pdf.go
 */
package pdf

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"

	"chunDoc/system/core/log"
	"chunDoc/system/util/file"
	"chunDoc/system/util/sys"
)

//尺寸
type PdfSizeTpye string

//横向纵向
type PdfOriTpye string

const (
	PdfPageSizeA4           PdfSizeTpye = "A4"
	PdfPageSizeA3           PdfSizeTpye = "A3"
	PdfPageSizeA5           PdfSizeTpye = "A5"
	PdfPageSizeLetter       PdfSizeTpye = "Letter"
	PdfOrinetationPortrait  PdfOriTpye  = "Portrait"  //纵向
	PdfOrinetationLandscape PdfOriTpye  = "Landscape" //横向
)

type PdfReuest struct {
	Url          string      `json:"url"` //url没有则取html数据生成
	Html         string      `json:"html"`
	Name         string      `json:"name" binding:"required"`
	Format       PdfSizeTpye `json:"format"`
	PageSize     PdfOriTpye  `json:"pageSize"`
	SavePath     string      `json:"SavePath"`
	WindowStatus string      `json:"WindowStatus"` //WindowStatus 表示等待状态开始执行
}

type PdfReturn struct {
	Name     string
	Url      string
	Fullpath string
}

type PdfUtil struct {
	HtmlTempdir string //临时html文件目录
	PdfDir      string //pdf文件目录
	Wkpdfpath   string //wkhtmltopdf路径
	HeaderPath  string //页眉路径
	FooterPath  string //页脚路径
	CoverPath   string //封面路径
	CssPath     string //css路径
}

//初始化
func Instance() *PdfUtil {
	basedir := sys.ExecutePath() + "\\" //根目录
	wk := basedir + "tools\\wkhtmltopdf.exe"

	htmltemp := basedir + "static\\pdf\\temp\\"
	file.IsNotExistMkDir(htmltemp)

	return &PdfUtil{
		HtmlTempdir: htmltemp,
		PdfDir:      basedir + "static\\pdf\\",
		Wkpdfpath:   wk,
		HeaderPath:  basedir + "static\\pdfconfig\\header.html",
		FooterPath:  basedir + "static\\pdfconfig\\footer.html",
		CoverPath:   basedir + "static\\pdfconfig\\cover.html",
		CssPath:     basedir + "static\\pdfconfig\\style.css",
	}
}

//创建临时html文件
func (thisClass *PdfUtil) CreateHtml(r *PdfReuest) (error, string) {
	htmlfile := thisClass.HtmlTempdir + r.Name + ".html"
	//读取css内容
	cssbytes, err := ioutil.ReadFile(thisClass.CssPath)
	var css string
	if err != nil {
		css = ""
	} else {
		css = fmt.Sprintf(`<style> %s
		</style>`, string(cssbytes))
	}

	//合并内容
	html := "<html><head><meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /></head><body>" + css + r.Html + "</body></html>"
	err = ioutil.WriteFile(htmlfile, []byte(html), 0644)
	return err, r.Name + ".html"

}

//创建pdf文件
func (thisClass *PdfUtil) CreatePdf(r *PdfReuest) (error, PdfReturn) {
	//时间戳命名
	//	r.Name = fmt.Sprintf("%s_%d", r.Name, time.Now().UnixMicro())
	//创建临时html文件
	err, name := thisClass.CreateHtml(r)
	if err != nil {
		return err, PdfReturn{}
	}
	//默认A4
	if r.Format == "" {
		r.Format = PdfPageSizeA4
	}

	//默认纵向
	if r.PageSize == "" {
		r.PageSize = PdfOrinetationPortrait
	}

	args := []string{"--page-size", string(r.Format), "--orientation", string(r.PageSize), "--title", r.Name, "--header-spacing", "5", "--footer-spacing", "5", "--header-html", thisClass.HeaderPath, "--footer-html", thisClass.FooterPath}

	if r.WindowStatus != "" {
		args = append(args, "--window-status", r.WindowStatus)
	}

	if !file.CheckNotExist(thisClass.CoverPath) {
		args = append(args, "cover", thisClass.CoverPath)
	}

	//输出位置
	actPdfPATH := thisClass.PdfDir
	if r.SavePath != "" {
		actPdfPATH = r.SavePath + "/"
	}
	fullpath := fmt.Sprintf("%s%s%s", actPdfPATH, r.Name, ".pdf")

	//要生成的url
	TogenUrl := thisClass.HtmlTempdir + name
	if r.Url != "" {
		TogenUrl = r.Url
	}
	args = append(args, TogenUrl, fullpath)

	err = Exec(thisClass.Wkpdfpath, args...)
	if err != nil {
		return err, PdfReturn{}
	}
	return nil, PdfReturn{
		Name:     r.Name,
		Url:      fmt.Sprintf("/pdf/%s.pdf", r.Name),
		Fullpath: fullpath,
	}
}

//执行命令
func Exec(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		log.Write(log.Error, fmt.Sprintln("exec the cmd ", name, " failed"))
		return err
	}
	// 正常⽇志
	logScan := bufio.NewScanner(stdout)
	go func() {
		text := ""
		for logScan.Scan() {
			text = text + logScan.Text() + "\n"
		}
		log.Write(log.Error, text)
	}()
	// 错误⽇志
	errBuf := bytes.NewBufferString("")
	scan := bufio.NewScanner(stderr)
	for scan.Scan() {
		s := scan.Text()
		errBuf.WriteString(s)
		errBuf.WriteString("\n")
	}
	// 等待命令执⾏完
	cmd.Wait()
	if !cmd.ProcessState.Success() {
		// 执⾏失败，返回错误信息
		return errors.New(errBuf.String())
	}
	return nil
}

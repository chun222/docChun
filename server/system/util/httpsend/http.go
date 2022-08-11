/*
 * @Date: 2022-06-08 09:02:59
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-06-16 12:14:18
 * @FilePath: \server\system\util\httpsend\http.go
 */
package httpsend

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"chunDoc/system/core/log"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) (error, string) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Write(log.Error, err.Error())
			return err, ""
		}
	}
	return nil, result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data []byte, contentType string) (error, []byte) {
	if contentType == "" {
		contentType = "application/json"
	}
	// 超时时间：60秒
	client := &http.Client{Timeout: 4 * 60 * time.Second}
	resp, err := client.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		log.Write(log.Error, err.Error())
		return err, []byte{}
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Write(log.Error, err.Error())
		return err, []byte{}
	}
	return nil, result // 返回json字符串
}

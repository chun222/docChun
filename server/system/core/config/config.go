/*
 * @Date: 2022-02-14 12:00:49
 * @LastEditors: 春贰
 * @Desc:
 * @LastEditTime: 2022-07-15 09:43:28
 * @FilePath: \server\system\core\config\config.go
 */
package config

import (
	"io/ioutil" //
	"log"

	"chunDoc/system/common/initial"
	"chunDoc/system/util/file"
	"github.com/spf13/viper" //配置
)

var c *conf

func Instance() *conf {
	if c == nil {
		InitConfig("./config.toml")
	}
	return c
}

type conf struct {
	DB     DBConf
	App    AppConf
	Zaplog ZapLogConf
	Jwt    JwtConf
	Upload UploadConf
	Pdf    PdfConf
}

type DBConf struct {
	DBType     string
	DBUser     string
	DBPwd      string
	DBHost     string
	DBName     string
	DBloglevel string
	DBoutpath  string
}

type AppConf struct {
	Key      string
	HttpPort int `json:"http-port"`
	WsPort   int
	Env      string
	PageSize int
}

type JwtConf struct {
	JwtSecret   string
	BufferTime  int64
	ExpiresTime int64
	Issuer      string
}

type ZapLogConf struct {
	Director string ` json:"director"`
	Level    string ` json:"level" `
}

type UploadConf struct {
	Path      string   ` json:"path"`
	MaxSize   float64  ` json:"maxsize" `
	AllowType []string ` json:"allowtype" `
}

type PdfConf struct {
	Path string ` json:"path"`
}

func InitConfig(tomlPath ...string) {
	if len(tomlPath) > 1 {
		log.Fatal("配置路径数量不正确")
	}
	if file.CheckNotExist(tomlPath[0]) {
		err := ioutil.WriteFile(tomlPath[0], []byte(initial.ConfigToml), 0777)
		if err != nil {
			log.Fatal("无法写入配置模板", err.Error())
		}
		log.Fatal("配置文件不存在，已在根目录下生成示例模板文件【config.toml】，请修改后重新启动程序！")
	}
	v := viper.New()
	v.SetConfigFile(tomlPath[0])
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("配置文件读取失败: ", err.Error())
	}
	err = v.Unmarshal(&c)
	if err != nil {
		log.Fatal("配置解析失败:", err.Error())
	}
}

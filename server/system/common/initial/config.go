/*
 * @Date: 2022-02-14 14:09:57
 * @LastEditors: 春贰
 * @Desc:
 * @LastEditTime: 2022-08-11 17:18:54
 * @FilePath: \server\system\common\initial\config.go
 */
package initial

const DataLang = `[{"name":"中文","dir":"zh-cn"},{"name":"English","dir":"en-us"}]`

const ConfigToml = `[app]
Key = "ADSDWW1DSADSADSAWJJK"   
HttpPort = 8023
WsPort = 2345   ###websocket 端口
PageSize = 20
Env = "dev"   ####运行环境 dev/prod

# 数据库配置 
# mysql
[db]
DBType = "mysql"
DBName = "tjs_ems_v2"
DBUser = "root"
DBPwd = "123456"
DBHost = "127.0.0.1:3306"
DBloglevel="info"        #日志级别 close/info/warn/error
DBoutpath = "databackup"

 
# 日志配置
# level:debug/info/warn/error/panic/fatal 
[zaplog]
director = 'runtime/log'
level = 'debug'

#Jwt 
[jwt]
JwtSecret = "ADSDWW#DSADSADSAWJJK"
BufferTime = 8640     #单位秒,缓冲时间
ExpiresTime = 8640000    #单位秒,过期时间
Issuer     = "tjs_ems_v2"

[upload]
Path = 'upload'   ###首尾不带"/"
MaxSize = 500   ##单位：M
AllowType = ["png", "jpeg", "jpg", "zip"]  

[pdf]
Path = 'runtime/pdf' 
`

const ZhMd = `# 中文
[{"name":"中文","dir":"zh-cn"},{"name":"English","dir":"en-us"}]`

const EnMd = `# EN`

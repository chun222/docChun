/*
 * @Date: 2022-02-14 14:09:57
 * @LastEditors: 春贰
 * @Desc:
 * @LastEditTime: 2022-08-19 17:14:29
 * @FilePath: \server\system\common\initial\config.go
 */
package initial

const DataLang = `[{"name":"中文","dir":"zh-cn"},{"name":"English","dir":"en-us"}]`

const DataVersion = `[{"name":"v01","dir":"v01"},{"name":"vTest","dir":"v02"}]`

const ConfigToml = `[app]
Key = "ADSDWW1DSADSADSAWJJK"   
HttpPort = 8024
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

const ZhMd = ` position:1
:info:

info 
**ssss**

ss
:: 
:success: 
success 
::
:warning: 
warning 

::
:error: 
error 
::  
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题 

*斜体文本*
_斜体文本_
**粗体文本**
__粗体文本__
***粗斜体文本***
___粗斜体文本___ 

***

* * *

*****

- - - 

----------

~~删除线~~

<u>下划线</u>

* 第一行
* 第二行

- 第一行
- 第二行

+ 第一行
+ 第二行

1.第一行
  - 第一点
  - 第二点

 

> 一级引用
> > 二级引用
> > > 三级引用

> 区块中使用列表
> 1. 第一项
> 2. 第二项
>> + 第一项
>> + 第二项
>> + 第三项

[这是百度](https://www.baidu.com)

 ###### 哈哈哈 
 

> #### 这是一个四级标题 


> 1. 这是第一行列表项
> 2. 这是第二行列表项

:success:

abv

::

:error:

**ssss**

### 事实上

Topic 1:  D
escription 1   
:Topic 1:  Description 1   

::   
|标题|标题|标题|
|:---|:---:|---:|
|居左测试文本|居中测试文本|居右测试文本|
|居左测试文本1|居中测试文本2|居右测试文本3|
|居左测试文本11|居中测试文本22|居右测试文本33|
|居左测试文本111|居中测试文本222|居右测试文本333|

![中文](./docimg/5.png)

![中文](http://localhost:8082/src/assets/vue.svg)
 

## 混合 
:warning: 
Topic 1:  Description 1   
:Topic 1:  Description 1  
::
$$x = {-b \pm \sqrt{b^2-4ac} \over 2a}.$$
# 事实上
啊`

const EnMd = `# EN`

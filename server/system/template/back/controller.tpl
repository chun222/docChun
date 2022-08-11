/* 
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 自动生成
 */
package {{.TableName}}

import (
	"chunDoc/system/core/request"
	"chunDoc/system/core/response"
	"chunDoc/system/model/CommonModel" 
	 AppDbModel "chunDoc/app/model/DbModel"
	 AppRequestModel "chunDoc/app/model/RequestModel"
	 {{.StructName}}Service "chunDoc/app/service" 
	"github.com/gin-gonic/gin"
)

var {{.StructName}}Ser {{.StructName}}Service.{{.StructName}}Service

/**
*
* @api {post} /{{.TableName}}/list {{.StructName}}
* @apiName {{.StructName}}
* @apiGroup {{.StructName}}
* @apiVersion  1.0.0
*
* @apiBody  {String}  page
* @apiBody  {String}  pageSize
* @apiSampleRequest /user/list
* @apiHeader {String} Authorization 唯一token
*
* @apiParamExample  {json} Request-Example:
* {
*    key : "value",
* }
*
*
* @apiSuccessExample {json} Success-Response:
*
{
 "code": 0,
	"data": {
	},
	"msg": "sucess"
}
*
*
*/
func List(c *gin.Context) {
	var r AppRequestModel.{{.StructName}}
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}

	if err, list, total := {{.StructName}}Ser.List(r); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(
			CommonModel.PageResult{
				List:     list,
				Total:    total,
				Page:     r.Page,
				PageSize: r.PageSize,
			}, c)
	}

}

/**
*
* @api {post} /{{.TableName}}/add 新增
* @apiName 新增{{.TableName}}
* @apiGroup {{.TableName}}
* @apiVersion  1.0.0
* @apiHeader {String} Authorization 唯一token
* @apiBody  {String}  apiBody
* @apiSampleRequest /{{.TableName}}/add
*
* @apiParamExample  {json} Request-Example:
*{"login_name":"12s1","User_Name":"asss","password":"ad","email":"sss@qq.com"}
*
*
* @apiSuccessExample {json} Success-Response:
*
{
    "code": 0,
    "data": 12,  //用户id
    "msg": "success"
}
*
*
*/
func Add(c *gin.Context) {
	var r AppDbModel.{{.StructName}}
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	} 
	if err = {{.StructName}}Ser.Add(&r); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(r.Id, c)
	}
}

/**
*
* @api {post} /{{.TableName}}/edit 编辑
* @apiName 编辑{{.TableName}}
* @apiGroup {{.TableName}}
* @apiVersion  1.0.0
* @apiHeader {String} Authorization 唯一token
* @apiBody  {String}  apiBody
* @apiSampleRequest /{{.TableName}}/edit
*
* @apiParamExample  {json} Request-Example:
*{"login_name":"12s1","User_Name":"asss","password":"ad","email":"sss@qq.com"}
*
*
* @apiSuccessExample {json} Success-Response:
*
{
    "code": 0,
    "data": 12,  //用户id
    "msg": "success"
}
*
*
*/
func Edit(c *gin.Context) {
	var r AppDbModel.{{.StructName}}
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}

	if err = {{.StructName}}Ser.Edit(&r); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

/**
*
* @api {post} /{{.TableName}}/romove 删除
* @apiName 删除{{.TableName}}
* @apiGroup {{.TableName}}
* @apiVersion  1.0.0
* @apiHeader {String} Authorization 唯一token
* @apiBody  {String}  apiBody
* @apiSampleRequest /{{.TableName}}/romove
*
* @apiParamExample  {json} Request-Example:
*{"login_name":"12s1","User_Name":"asss","password":"ad","email":"sss@qq.com"}
*
*
* @apiSuccessExample {json} Success-Response:
*
{
    "code": 0,
    "data": 12,  //用户id
    "msg": "success"
}
*
*
*/
func Romove(c *gin.Context) {
	var r CommonModel.Ids
	err, msg := request.Binding(&r, c)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}

	if err = {{.StructName}}Ser.Remove(r); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}
 
/* 
 * @Desc: 自动生成
 * @LastEditTime: 2022-03-21 11:07:54
 * @FilePath: \server\system\template\front\api.tpl
 */
import http from '../http'

const Api = {
    list: '/{{.TableName}}/list',
    add: '/{{.TableName}}/add',
    edit: '/{{.TableName}}/edit', 
    remove: '/{{.TableName}}/remove', 
}

export const list = (data) => {   
    return http.request({
        url: Api.list,
        data: data,
        method: 'POST'
    })
} 
export const add = data => {
    return http.request({
        url: Api.add,
        data: data,
        method: 'POST'
    })
}

export const edit = data => { 
    return http.request({
        url: Api.edit,
        data: data,
        method: 'POST'
    })
}

export const remove = data => {
    return http.request({
        url: Api.remove,
        data: data,
        method: 'POST'
    })
}

 
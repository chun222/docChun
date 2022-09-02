/*
 * @Date: 2022-03-04 14:27:46
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-09-02 16:09:38
 * @FilePath: \pages\src\api\module\base.ts
 */

import http from '../http'
import { TypeDoclist, TypePath, TypeSearch, AliasDirType } from '@/model';

//获取文档列表

export const doclist = (data: TypeDoclist) => {
  return http.request({
    url: '/doclist',
    data: data,
    method: 'post'
  })
}



export const read = (data: TypePath) => {
  return http.request<string>({
    url: '/read',
    data: data,
    method: 'post'
  })
}

export const savecontent = (data: TypePath) => {
  return http.request<string>({
    url: '/savecontent',
    data: data,
    method: 'post'
  })
}




export const search = (data: TypeSearch) => {
  return http.request({
    url: '/search',
    data: data,
    method: 'post'
  })
}

//所有配置
export const allconfig = () => {
  return http.request({
    url: '/allconfig',
    method: 'post'
  })
}

//保存配置 
export const saveconfigs = (data: { projects: AliasDirType[]; versions: AliasDirType[], langs: AliasDirType[] }) => {
  return http.request({
    url: '/saveconfigs',
    data: data,
    method: 'post'
  })
}

//修改或新增文件菜单
export const create_update_file = (data: any) => {
  return http.request({
    url: '/create_update_file',
    data: data,
    method: 'post'
  })
}

export const removefile = (data: { path: string }) => {
  return http.request({
    url: '/removefile',
    data: data,
    method: 'post'
  })
}


//上传文件
export const upload = (data: { file: File }) => {
  return http.request({
    headers:{
      'Content-Type': 'multipart/form-data',
    },
    url: '/uploadfile',
    data: data,
    method: 'post'
  })
}



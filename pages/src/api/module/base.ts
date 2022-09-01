/*
 * @Date: 2022-03-04 14:27:46
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-09-01 12:11:02
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


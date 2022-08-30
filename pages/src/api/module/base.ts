/*
 * @Date: 2022-03-04 14:27:46
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-30 14:05:43
 * @FilePath: \pages\src\api\module\base.ts
 */

import http from '../http' 
import axios, { Method, AxiosInstance, AxiosRequestConfig, AxiosPromise, AxiosInterceptorManager, AxiosResponse } from 'axios';

 

//获取文档列表
type TypeDoclist = {
  project:string;
  lang:string;
  version:string ;
} 
export const doclist = (data: TypeDoclist) => {
  return http.request({
    url: '/doclist',
    data: data,
    method: 'post'
  })
}


type TypePath =  {
  project:string;
  lang:string;
  version:string ;
  path:string
}

export const read = (data: TypePath) => {
  return http.request<string>({
    url: '/read',
    data: data,
    method: 'post'
  })
}


type TypeSearch =  {
  project:string;
  lang:string ; 
  version:string ;
  keyword:string;
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

 
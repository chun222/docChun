/*
 * @Date: 2022-03-04 14:27:46
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-14 01:03:32
 * @FilePath: \pages\src\api\module\base.ts
 */

import http from '../http' 
import axios, { Method, AxiosInstance, AxiosRequestConfig, AxiosPromise, AxiosInterceptorManager, AxiosResponse } from 'axios';

 

//获取文档列表
interface TypedDoclist {
  langname:string
}
export const doclist = (data: TypedDoclist) => {
  return http.request({
    url: '/doclist',
    data: data,
    method: 'post'
  })
}
 
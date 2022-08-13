/*
 * @Date: 2022-08-14 00:29:26
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-14 00:34:21
 * @FilePath: \pages\src\api\type.ts
 */ 

import type { AxiosRequestConfig, AxiosResponse } from 'axios'

export interface RequestInterceptors {
  requestIntercetor?: (config: AxiosRequestConfig) => AxiosRequestConfig
  requestIntercetorCatch?: (error: any) => any
  responseInterceptor?: (res: AxiosResponse) => AxiosResponse
  responseInterceptorCatch?: (error: any) => any
}

export interface RequestConfig extends AxiosRequestConfig {
  interceptors?: RequestInterceptors
}

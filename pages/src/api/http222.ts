/*
 * @Date: 2022-08-14 00:28:54
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-14 00:34:07
 * @FilePath: \pages\src\api\http.ts
 */
import axios from 'axios'
import type { AxiosInstance, AxiosResponse } from 'axios'
import type { RequestInterceptors, RequestConfig } from './type'

class Request {
  // axios 实例
  instance: AxiosInstance
  // 拦截器对象
  interceptorsObj?: RequestInterceptors
  // 使用拦截器
  constructor(config: RequestConfig) {
    this.instance = axios.create(config)

    // 1. 使用实例拦截器
    this.interceptorsObj = config.interceptors

    this.instance.interceptors.request.use(
      this.interceptorsObj?.requestIntercetor,
      this.interceptorsObj?.requestIntercetorCatch
    )
    this.instance.interceptors.response.use(
      this.interceptorsObj?.responseInterceptor,
      this.interceptorsObj?.responseInterceptorCatch
    )

    // 2. 全局拦截器(所有实例都有的拦截器)
    this.instance.interceptors.request.use(
      (config: RequestConfig) => {
        console.log('所有的实例都有的拦截器: 请求成功拦截')
        return config
      },
      //  请求失败拦截
      (err: any) => err
    )
    this.instance.interceptors.response.use(
      (res: AxiosResponse) => {
        console.log('所有实例都有的拦截器：响应成功拦截')
        return res
      },
      // 响应失败拦截
      (err: any) => err
    )
  }
  // 3. 为单个请求添加拦截器
  request(config: RequestConfig): void {
    // 如果我们为单个请求添加拦截器，这里使用单个请求的拦截
    if (config.interceptors?.requestIntercetor) {
      config = config.interceptors.requestIntercetor(config)
    }
    this.instance.request(config).then((res) => {
      // 如果我们为单个响应添加拦截器，这里使用单个响应的拦截
      if (config.interceptors?.responseInterceptor) {
        res = config.interceptors.responseInterceptor(res)
      }
      console.log(res)
    })
  }
}

export default Request 
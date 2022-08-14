/*
 * @Date: 2022-03-04 11:07:17
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-14 00:59:04
 * @FilePath: \pages\src\api\http.ts
 */

import axios, { Method, AxiosInstance, AxiosRequestConfig, AxiosPromise, AxiosInterceptorManager, AxiosResponse } from 'axios';

import { notification, message as Msg } from "ant-design-vue";



const refreshToken = (response: AxiosResponse) => {
  let token = response.headers.authorization
  if (token) {
    localStorage.setItem('USER_TOKEN', token);
  }
}

notification.config({
  duration: 5
})

const baseurl = "http://localhost:8023/";


type ResultDataType = {
  code: number;
  data?: any;
  msg?: string;
}

interface NewAxiosInstance extends AxiosInstance {
  /* 
  设置泛型T，默认为any，将请求后的结果返回变成AxiosPromise<T>
  */
  <T = any>(config: AxiosRequestConfig): AxiosPromise<ResultDataType>;
  interceptors: {
    request: AxiosInterceptorManager<AxiosRequestConfig>;
    response: AxiosInterceptorManager<AxiosResponse<ResultDataType>>;
  }
}



interface RequestInterceptors {
  // 请求拦截
  requestInterceptors?: (config: AxiosRequestConfig) => AxiosRequestConfig
  requestInterceptorsCatch?: (err: any) => any
  // 响应拦截
  responseInterceptors?: (config: AxiosResponse) => AxiosResponse
  responseInterceptorsCatch?: (err: any) => any
}
// 自定义传入的参数
interface RequestConfig extends AxiosRequestConfig {
  interceptors?: RequestInterceptors
}

class Http { 
  // axios 实例
  instance: AxiosInstance; 
  // 拦截器对象
  interceptorsObj?: RequestInterceptors
  config: AxiosRequestConfig<any>;
  constructor(config: AxiosRequestConfig) {
    this.config = config || {
      timeout: 1000 * 60 * 3, //默认三分钟
      withCredentials: true,
      baseURL: baseurl,
      headers: {
        "Content-Type": "application/json; charset=utf-8",
      }
    };
  }

  interceptors(instance: NewAxiosInstance) {
    /**
     * 请求拦截器
     */
    instance.interceptors.request.use(
      config => {
        const token = localStorage.getItem("USER_TOKEN");
        if (token) {
          config.headers["Authorization"] = token;
        }

        // config.cancelToken = new axios.CancelToken(async cancel => {
        //   await store.dispatch("app/execCancelToken", { cancelToken: cancel });
        // });
        return config;
      },
      error => {
        return Promise.reject(error);
      }
    );
    /** 响应拦截 */
    instance.interceptors.response.use(
      response => {
        refreshToken(response);//有token 就刷新 
        if (response.data.code > 1) {
          notification.error({
            message: response.data.code,
            description: response.data.msg
          });
        }
        // return Promise.resolve;
        return response.data;
      },
      error => {
        console.log('error', error.response);
        if (error.response) {
          refreshToken(error.response); //有token 就刷新 
          if (error.response.status === 403) {
            notification.error({
              message: "permision error",
              description: error.response.request.responseURL
            });
          } else if (error.response.status === 401) {
            // store.dispatch("user/logout").then(() => {
            // });
          } else if (error.response.status === 404) {
            notification.error({
              message: "not find api",
              description: "not find api"
            });
          }
        } else {
          let { message } = error;
          Msg.error(message);
        }
        return Promise.reject("error response");
      }
    );
  } 
  // request<T>(config: RequestConfig): Promise<T> {
  //   return new Promise((resolve, reject) => {
  //     // 如果我们为单个请求设置拦截器，这里使用单个请求的拦截器
  //     if (config.interceptors?.requestInterceptors) {
  //       config = config.interceptors.requestInterceptors(config)
  //     }
  //     this.instance
  //       .request<any, ResultDataType>(config)
  //       .then(res => {
  //         // 如果我们为单个响应设置拦截器，这里使用单个响应的拦截器
  //         if (config.interceptors?.responseInterceptors) {
  //           res = config.interceptors.responseInterceptors<T>(res)
  //         }

  //         resolve(res)
  //       })
  //       .catch((err: any) => {
  //         reject(err)
  //       })
  //   })
  // }

  request(options: AxiosRequestConfig){
    const instance:NewAxiosInstance = axios.create();
    const requestOptions = Object.assign({}, this.config, options);
    this.interceptors(instance);
    return instance(requestOptions);
  }
}


const http = new Http({});
export default http;

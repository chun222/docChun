/*
 * @Date: 2022-03-04 11:07:17
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-11 17:52:29
 * @FilePath: \web\src\api\http.js
 */
 
import axios from "axios";
import { notification, message as Msg } from "ant-design-vue"; 

const refreshToken = (response)=>{
  let token = response.headers.authorization 
  if (token) { 
      localStorage.setItem('USER_TOKEN', token);
  }
}

notification.config({
  duration : 5
})

const baseurl = "http://localhost:8023/";

class Http {

  constructor(config) {
    this.config = config || {
      timeout: 1000 * 60 * 3, //默认三分钟
      withCredentials: true,
      baseURL: baseurl,
      headers: {
        "Content-Type": "application/json; charset=utf-8",
      }
    };
  }

  interceptors(instance) {
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
        if (response.data.code>1) {
          notification.error({
            message: response.data.code,
            description: response.data.msg
          });
        }
        return response.data;
      },
      error => {
        console.log('error',error.response);
        if (error.response) {
          refreshToken(error.response); //有token 就刷新 
          if (error.response.status === 403) {
            notification.error({
              message: "permision error",
              description: error.response.request.responseURL
            });
          }else if(error.response.status === 401) {   
             store.dispatch("user/logout").then(() => {  
            });
          }else if(error.response.status === 404){
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

  request(options) {
    const instance = axios.create();
    const requestOptions = Object.assign({}, this.config, options);
    this.interceptors(instance); 
    return instance(requestOptions);
  }
}

const http = new Http();
export default http;

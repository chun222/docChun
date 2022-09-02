import axios from "axios";
import { notification, message as Msg } from "ant-design-vue";
import type { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
type Result<T> = {
  code: number;
  msg: string;
  data: T;
};

const refreshToken = (response: AxiosResponse) => { 
  let token = response.headers.authorization
  if (token) {
    localStorage.setItem('USER_TOKEN', token);
  }
}

notification.config({
  duration: 5
})


const baseurl = import.meta.env.VITE_API_URL;
console.log(baseurl,"baseurl");
class Http {
  // axios 实例
  instance: AxiosInstance;
  // 基础配置，url和超时时间
  baseConfig: AxiosRequestConfig = {
    baseURL: baseurl, timeout: 60000, withCredentials: true,
    headers: {
      "Content-Type": "application/json; charset=utf-8",
    }
  };

  constructor(config: AxiosRequestConfig) {
    // 使用axios.create创建axios实例
    this.instance = axios.create(Object.assign(this.baseConfig, config));

    this.instance.interceptors.request.use(
      (config: AxiosRequestConfig) => {
        // 一般会请求拦截里面加token
        const token = localStorage.getItem("USER_TOKEN");
        if (token) {
          config.headers["Authorization"] = token;
        }
        return config;
      },
      (err: any) => { 
        return Promise.reject(err);
      }
    );

    this.instance.interceptors.response.use(
      <T>(res: AxiosResponse<Result<T>>) => {
        refreshToken(res);//有token 就刷新 
        // 直接返回res，当然你也可以只返回res.data
        return res.data;
      },
      (err: any) => {
        console.log(err);
        // 这里用来处理http常见错误，进行全局提示
        refreshToken(err.response);//有token 就刷新 
        let message = "";
        switch (err.response.status) {
          case 400:
            message = "请求错误(400)";
            break;
          case 401:
            message = "未授权，请重新登录(401)";
            // 这里可以做清空storage并跳转到登录页的操作
            break;
          case 403:
            message = "拒绝访问(403)";
            break;
          case 404:
            message = "请求出错(404)";
            break;
          case 408:
            message = "请求超时(408)";
            break;
          case 500:
            message = "服务器错误(500)";
            break;
          case 501:
            message = "服务未实现(501)";
            break;
          case 502:
            message = "网络错误(502)";
            break;
          case 503:
            message = "服务不可用(503)";
            break;
          case 504:
            message = "网络超时(504)";
            break;
          case 505:
            message = "HTTP版本不受支持(505)";
            break;
          default:
            message = `连接出错(${err.response.status})!`;
        }
        // 这里错误消息可以使用全局弹框展示出来 
        notification.error({
          message: err.response.status,
          description: message
        });

        // 这里是AxiosError类型，所以一般我们只reject我们需要的响应即可
        return Promise.reject(err.response);
      }
    );
  }

  // 定义请求方法
  public request<T = any>(config: AxiosRequestConfig): Promise<Result<T>> {
    //合并配置
    return this.instance.request(Object.assign({}, this.baseConfig, config));
  }

  public get<T = any>(
    url: string,
    config?: AxiosRequestConfig
  ): Promise<Result<T>> {
    return this.instance.get(url, Object.assign({}, this.baseConfig, config));
  }
  public post<T = any>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<Result<T>> {
    return this.instance.post(url, data, Object.assign({}, this.baseConfig, config));
  }
  public put<T = any>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<Result<T>> {
    return this.instance.put(url, data, Object.assign({}, this.baseConfig, config));
  }
  //Promise<AxiosResponse<Result<T>>>
  public delete<T = any>(
    url: string,
    config?: AxiosRequestConfig
  ): Promise<Result<T>> {
    return this.instance.delete(url, Object.assign({}, this.baseConfig, config));
  }
}
const http = new Http({});
export default http;
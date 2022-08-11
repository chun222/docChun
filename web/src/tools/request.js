/*
 * @Date: 2021-10-21 08:36:40
 * @LastEditors: 邓春宜
 * @Desc: 
 * @LastEditTime: 2021-11-29 11:34:34
 * @FilePath: \tianjushi-kanban\src\util\request.js
 */
import axios from 'axios';
export function request(config) { 
    
    const instance = axios.create({
        baseURL:process.env.NODE_ENV === 'production' ? 'http://120.27.240.124:2225' :  'http://localhost:2225',
        timeout: 5000
    }) 
    // 添加请求拦截器
    instance.interceptors.request.use(config => {
        // loadingInstance = Loading.service({
        //     lock: true,
        //     text: 'loading...'
        // })
       // console.log(config);
        return config ;
    })
    // 添加响应拦截器
    instance.interceptors.response.use(response => {
        // loadingInstance.close() 
         if (response.status===200) {
            return response.data; 
         }else{
            return response; 
         }
         
    }, error => {
        console.log('TCL: error', error)
        const msg = error.Message !== undefined ? error.Message : ''
        // Message({
        //   message: '网络错误' + msg,
        //   type: 'error',
        //   duration: 3 * 1000
        // })
        // loadingInstance.close()
        return Promise.reject(error)
    }) 
    return instance(config); 
}

 
//获取单个
export function requestPoint(KanBanName,fields) {   
   return request({
       url:"/api/data/point",
       method: 'post',
       data: {
        name: KanBanName,
        field: fields
      }
   });
}

//获取列表
export function requestList(KanBanName,fields,count) {  
    return request({
        url:"/api/data/list",
        method: 'post',
        data: {
         name: KanBanName,
         field: fields,
         count:count
       }
    });
 }

 
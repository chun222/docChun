/*
 * @Date: 2022-03-04 14:27:46
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-11 17:49:35
 * @FilePath: \web\src\api\module\BaseLogin.js
 */

import http from '../http' 

/// 登录
export const doclist = data => {
  return http.request({
    url: '/doclist',
    data: data,
    method: 'post'
  })
}
 
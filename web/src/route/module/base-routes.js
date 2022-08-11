/*
 * @Date: 2022-08-10 08:56:08
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-11 09:21:07
 * @FilePath: \web\src\route\module\base-routes.js
 */
export default [
  {
    path: '/',
    component: () => import('@/view/index.vue'),
    meta: {
      title: '首页'
    },
    hidden: 0,
  }, 
  {
    path: '/error/404',
    component: () => import('@/view/404.vue'),
    meta: {
      title: '404'
    },
  }
]
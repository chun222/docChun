/*
 * @Date: 2022-08-10 08:56:08
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-11 14:48:37
 * @FilePath: \web\src\route\module\base-routes.js
 */
export default [
  {
    name:'index',
    path: '/:id',
    component: () => import('@/view/index/index.vue'),
    meta: {
      title: '首页'
    },
    hidden: 0,
  }, 
  {
    name:'error/404',
    path: '/error/404',
    component: () => import('@/view/404.vue'),
    meta: {
      title: '404'
    },
  }
]
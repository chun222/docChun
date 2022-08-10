/*
 * @Date: 2021-11-30 15:42:07
 * @LastEditors: 春贰
 * @Desc: 
 * @LastEditTime: 2022-08-10 09:58:04
 * @FilePath: \web\src\route\module\base-routes.js
 */
/// 基础路由  

export default [  
  {
    path: '/',
    redirect: "/dashboard/console",
    hidden: 0,
  },  
   {
    path: '/login',
    component: () => import('@/view/base/login.vue'),
    meta: {
      title: 'login', 
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

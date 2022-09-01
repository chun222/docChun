/*
 * @Date: 2022-08-10 08:56:08
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-09-01 08:44:05
 * @FilePath: \pages\src\route\module\base-routes.ts
 */
export default [
  {
    name:'/',
    path: '/:project?/:version?/:lang?/:page?/:id?',
    component: () => import('@/view/index/index.vue'),
    meta: {
      title: '首页'
    },
    hidden: 0,
  }, 
  {
    name:'error_404',
    path: '/error/404',
    component: () => import('@/view/404.vue'),
    meta: {
      title: '404'
    },
  },
  {
    name:'admin',
    path: '/admin/:project?/:version?/:lang?/:page?',
    component: () => import('@/view/admin/admin.vue'),
    meta: {
      title: '后台'
    },
  }
]

/*
 * @Date: 2022-08-10 08:56:08
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-10 12:15:43
 * @FilePath: \web\src\route\index.js
 */
import { createRouter, createWebHashHistory } from 'vue-router'
import routes from './module/base-routes'
import NProgress from "nprogress";
import "nprogress/nprogress.css"; 

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

const setDocumentTitle = title => {
  document.title = `chunDoc-${title}`
}

//添加动态路由
router.beforeEach(to => {
  NProgress.start();
  const { meta } = to
  setDocumentTitle(meta.title)
  if (!router.hasRoute("admin")) {
    router.addRoute({ path: '/admin', name: 'admin', component: () => import('@/view/admin.vue') })
    return to.fullPath
  }
})

//控制权限
router.beforeEach((to, from, next) => {
  NProgress.start();
  if (!to.fullPath.includes('login') && !localStorage.getItem('USER_TOKEN') && false) {
    next({ path: '/login' }) 
  } else {
    if (router.getRoutes().map(it => it.path).includes(to.path)) {
      next()
    } else {
      next('/error/404')
    }
  }
})
 
router.afterEach((to, from) => {
  NProgress.done();
})

export default router
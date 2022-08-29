/*
 * @Date: 2022-08-10 08:56:08
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-29 11:07:10
 * @FilePath: \pages\src\route\index.ts
 */
import { createRouter, createWebHashHistory, RouteMeta } from 'vue-router'
import routes from './module/base-routes'
import NProgress from "nprogress";
import "nprogress/nprogress.css";

import { useStore } from "@/store/index";

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

const setDocumentTitle = (title: unknown) => {
  document.title = `chunDoc-${title}`
}

//添加动态路由
router.beforeEach(to => {
  const meta = to.meta
  setDocumentTitle(meta.title)
  //设置默认首页跳转
  // const store = useStore(); 
  // if (!router.hasRoute("/")) {
  //   NProgress.start(); 
  //   router.addRoute({ name: '/',  path:null, redirect:`/${store.project.dir}/${store.version.dir}/${store.lang.dir}`})
  //   return to.fullPath
  // }

  if (!router.hasRoute("admin")) {
    NProgress.start();
    router.addRoute({ path: '/admin', name: 'admin', component: () => import('@/view/admin.vue') })
    return to.fullPath
  }
})

//控制权限
router.beforeEach((to, from, next) => {
  //NProgress.start();
  if (!to.fullPath.includes('login') && !localStorage.getItem('USER_TOKEN') && false) {
    next({ path: '/login' })
  } else {
    // console.log("router",router.getRoutes(),to.path,to.fullPath,to.name);
    if (router.getRoutes().map(it => it.name).includes(to.name)) {
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
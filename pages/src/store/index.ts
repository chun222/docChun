/*
 * @Date: 2022-08-10 17:31:57
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 状态很少无需解构
 * @LastEditTime: 2022-08-29 11:22:37
 * @FilePath: \pages\src\store\index.ts
 */
import { defineStore } from 'pinia'


import { allconfig } from "@/api/module/base";

import router from "@/route/index" 



type AliasDirType = {
  name:string;
  dir:string;
} 
 
export const useStore = defineStore('main', {
  // 推荐使用 完整类型推断的箭头函数
  state: () => {
    return {
      // 所有这些属性都将自动推断其类型 
      theme: localStorage.getItem("theme")||'light',  //light or dark 
      //菜单
      menus: [],
      version:  {name:"",dir:""},
      lang: {name:"",dir:""},
      project: {name:"",dir:""},
      projects:[{name:"",dir:""}],
      versions:[{name:"",dir:""}],
      langs: [{name:"",dir:""}],
      routeLoading:false,
    }
  },
  getters: {
    //  doubleCount: (state) => state.counter * 2,
  },
  actions: {

    //初始化配置
    async InitConfig() {
      return allconfig().then((re) => {
        if (re.code == 0) { 
          this.versions = re.data.Version;
          this.langs = re.data.Lang;
          this.projects = re.data.Project;

          if (this.projects.length > 0) { 
            let cacheSet = router.currentRoute.value.params.project
            if(cacheSet === ""){
             this.changeProject(this.projects[0])
            }else{
              let finded =  this.projects.find((item:AliasDirType)=>item.dir==cacheSet)
              if(finded){
                this.changeProject(finded)
              }else{ 
                 this.changeProject(this.projects[0])
              }
            } 
          } 
          
          if (this.versions.length > 0) { 
            let cacheSet = router.currentRoute.value.params.version
            if(cacheSet == ""){
             this.changeVersion(this.versions[0])
            }else{
              let finded =  this.versions.find((item:AliasDirType)=>item.dir==cacheSet)
              if(finded){
                this.changeVersion(finded)
              }else{ 
                 this.changeVersion(this.versions[0])
              }
            } 
          } 

          if (this.langs.length > 0) { 
            let cacheSet =  router.currentRoute.value.params.lang
            if(cacheSet == ""){
             this.changeLang(this.langs[0])
            }else{
              let finded =  this.langs.find((item:AliasDirType)=>item.dir==cacheSet)
              if(finded){
                this.changeLang(finded)
              }else{ 
                 this.changeLang(this.langs[0])
              }
            } 
          }
          
        }
      })
    },
    //改变主题
    changeTheme(v: string) {
      this.theme = v;
      localStorage.setItem("theme", v)  
    },

     //改变项目
     changeProject(v: AliasDirType) {
      this.project = v;
    // localStorage.setItem("project", v.dir)
      
    },

    //改变版本
    changeVersion(v: AliasDirType,isChangeRouter?:boolean) {
      this.version = v;
    //  localStorage.setItem("version", v.dir)
    // if (isChangeRouter) {
    //     console.log("push==",Object.assign(router.currentRoute.value.params,{version: v.dir}));
    //  router.push({name:"/",  params:Object.assign(router.currentRoute.value.params,{version: v.dir})})
    // }
    },

    //改变语言
    changeLang(v: AliasDirType) {
      this.lang = v;
    //  localStorage.setItem("lang",  v.dir)
    },


  },
})
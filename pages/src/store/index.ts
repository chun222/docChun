/*
 * @Date: 2022-08-10 17:31:57
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 状态很少无需解构
 * @LastEditTime: 2022-08-13 22:06:33
 * @FilePath: \pages\src\store\index.ts
 */
import { defineStore } from 'pinia'



// useStore 可以是 useUser、useCart 之类的任何东西
// 第一个参数是应用程序中 store 的唯一 id
export const useStore = defineStore('main', {
    // 推荐使用 完整类型推断的箭头函数
    state: () => {
        return {
            // 所有这些属性都将自动推断其类型 
            theme: 'light',  //light or dark 
            //菜单
            menus:[],
        }
    },
    getters: {
      //  doubleCount: (state) => state.counter * 2,
    }, 
    actions: {
        //改变主题
        changeTheme(v:string){
            this.theme = v ;
        },
        // async getMenus(lang) {
        //     doclist({langname : lang}).then((re)=>{
        //       if (re.code==0) {
        //         this.menus = re.data ;
        //       }
        //     })
        // }, 
      },
})
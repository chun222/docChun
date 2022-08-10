/*
 * @Date: 2022-08-10 17:31:57
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 状态很少无需解构
 * @LastEditTime: 2022-08-10 17:45:44
 * @FilePath: \web\src\store\index.js
 */
import { defineStore } from 'pinia'

// useStore 可以是 useUser、useCart 之类的任何东西
// 第一个参数是应用程序中 store 的唯一 id
export const useStore = defineStore('main', {
    // 推荐使用 完整类型推断的箭头函数
    state: () => {
        return {
            // 所有这些属性都将自动推断其类型
            counter: 0,
            theme: 'dark',
            isAdmin: true,
        }
    },
    getters: {
        doubleCount: (state) => state.counter * 2,
    }, 
    actions: {
        //改变主题
        changeTheme(v){
            this.theme = v ;
        },
        async registerUser(login, password) {
          try {
            this.userData = await api.post({ login, password })
            showTooltip(`Welcome back ${this.userData.name}!`)
          } catch (error) {
            showTooltip(error)
            // 让表单组件显示错误
            return error
          }
        },
      },
})
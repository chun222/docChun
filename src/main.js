/*
 * @Date: 2022-08-08 14:16:39
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-10 09:07:02
 * @FilePath: \web\src\main.js
 */
import { createApp } from 'vue' 
import App from './App.vue'
import Router from "./route";

const app = createApp(App);
app.use(Router); 
app.mount('#app')

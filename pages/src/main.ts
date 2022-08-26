/*
 * @Date: 2022-08-08 14:16:39
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-26 11:24:23
 * @FilePath: \pages\src\main.ts
 */
import { createApp } from 'vue' 
import App from './App.vue'
import Router from "./route/index";

import registerComponents from './tools/registerComponents' 
import {  message } from 'ant-design-vue';
import 'ant-design-vue/dist/antd.variable.min.css';  
import './assets/style/root.css'; 
import './assets/style/main.css'; 
import './assets/style/theme/dark.css'; 
import './assets/style/theme/light.css';  
import { createPinia } from 'pinia'


const app = createApp(App);
app.config.globalProperties.$message = message;
app.use(createPinia()) 
//全局注册常用antd组件 
registerComponents.forEach((v)=>app.use(v))  
import { useStore } from "@/store/index";
const store = useStore();
//状态优先
store.InitConfig().then(()=>{
    app.use(Router); 
    app.mount('#app')
})




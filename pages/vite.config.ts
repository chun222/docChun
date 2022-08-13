/*
 * @Date: 2022-08-12 22:53:33
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-13 23:00:01
 * @FilePath: \pages\vite.config.ts
 */
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import path from "path";

export default defineConfig({
  plugins: [vue()],
  base: "./",  
  resolve: {
    alias: { 
      // 如果报错__dirname找不到，需要安装node,执行npm install @types/node --save-dev
        '@': path.resolve(__dirname, 'src'),
    },
  },
  build: {
    minify: 'terser',  //'terser'  npm add -D terser 
    chunkSizeWarningLimit: 1500, // 单个文件大小警告
    //清除console
    terserOptions: {
      compress: {
        drop_console: true,
        drop_debugger: true,
      },
    },
    rollupOptions: {
      // output:{
      //   manualChunks(id){ // 分包
      //     if(id.includes('node_modules')){
      //       return id.toString().split('node_modules/')[1].split('/')[0].toString();
      //     }
      //   }
      // }
    }
  }, 
  // 引入第三方的配置
  optimizeDeps: {
    include: [],
  },
});
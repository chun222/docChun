/*
 * @Date: 2022-07-23 11:03:25
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-10 12:19:12
 * @FilePath: \web\vite.config.js
 */

import { defineConfig, loadEnv } from 'vite';
import vue from '@vitejs/plugin-vue'

import path from 'path'

const config = ({ mode }) => {
  return defineConfig({
    plugins: [vue()],
    base: "./",
    server: {
      host: '0.0.0.0',
      port: loadEnv(mode, process.cwd()).VITE_PORT
    },
    resolve: {
      extensions: ['.js', '.vue', '.json'],
      alias: {
        "@": path.resolve(__dirname, "src"),
        "components": path.resolve(__dirname, "src/components"),
        "img": path.resolve(__dirname, "src/assets/image"),
        "images": path.resolve(__dirname, "src/assets/image"),
        "font": path.resolve(__dirname, "src/assets/font"),
      },
      // 打包配置
      build: {
        outDir: 'dist', // 指定输出路径
        assetsDir: 'assets', //指定生成静态资源的存放路径 
      },
    },
    // 开启less支持
    css: {
      preprocessorOptions: {
        less: {
          javascriptEnabled: true
        }
      }
    },
    hmr: {
      overlay: true
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
    }
  })
}

export default config
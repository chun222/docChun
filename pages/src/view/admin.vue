<!--
 * @Date: 2022-08-10 09:26:29
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-30 18:07:41
 * @FilePath: \pages\src\view\admin.vue
-->
<template>
  <div id="markdown" :style="{ height: mdHeight + 'px' }"></div>

  <!-- 曲线救国 标准方法获取不到 -->
  <div class="diyButtons">
    <span class="diyButtonspan" @click="saveMd">sssss</span>
  </div>
</template>

<script lang="ts">
import Cherry from "cherry-markdown";
import { CherryMarkdownConfig } from "@/tools/cherryMarkdown";
import "cherry-markdown/dist/cherry-markdown.css";

import { RouteParams, useRoute } from "vue-router";
import { useStore } from "@/store/index";

import { read } from "@/api/module/base";
//cherryInstance.export("pdf");
import { onMounted, ref, watch,h } from "vue"; 
import 'font-awesome/css/font-awesome.min.css'

import router from "@/route/index";
export default {
  setup() {

  
    const pageLoading = ref(true);

    const store = useStore();
    const route = useRoute();
     const BConfigClass =   new CherryMarkdownConfig({
      forceAppend: true,
    }).BasicConfig
   

    let cherryInstance: any;

    watch(
      () => route.params.page,
      (val: any) => {
        loaddocRaw(val as string);
      }
    );

    onMounted(() => {
      cherryInstance = new Cherry(BConfigClass);  

      document.querySelector(".cherry-toolbar").appendChild(document.querySelector(".diyButtons"))
      // document.querySelector(".diyButtons").childNodes.forEach((v)=>{ 
      //   document.querySelector(".cherry-toolbar").appendChild(v)
      // })
      // document.querySelector(".diyButtons").remove()
      
    });

     watch(
      () => store.InitOk,
      (v1: boolean,v2: boolean) => {
        const page = route.params.page as string;
          loaddocRaw(page);
      }
    );

    

    //请求页面数据
    const loaddocRaw = (page: string): void => {
      pageLoading.value = true;
      read({
        path: page,
        project: store.project.dir,
        version: store.version.dir,
        lang: store.lang.dir,
      }).then((v) => {
        if (v.code == 0) {
          pageLoading.value = false;
          if (v.data == "") {
            return;
          }
          cherryInstance.setMarkdown(v.data, false);
        }
      });
    };

    const saveMd = ()=>{
      console.log(cherryInstance.getMarkdown());
    }

    return {
      mdHeight: document.body.clientHeight - 135,
      saveMd
    };
  },
};
</script>

<style>
</style>
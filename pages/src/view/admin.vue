<!--
 * @Date: 2022-08-10 09:26:29
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-30 11:32:51
 * @FilePath: \pages\src\view\admin.vue
-->
<template>
  <div id="markdown" :style="{ height: mdHeight + 'px' }"></div>
</template>

<script lang="ts">
import Cherry from "cherry-markdown";
import { BasicConfig } from "@/tools/cherryMarkdown";
import "cherry-markdown/dist/cherry-markdown.css";

import { RouteParams, useRoute } from "vue-router";
import { useStore } from "@/store/index";

import { read } from "@/api/module/base";
//cherryInstance.export("pdf");
import { onMounted, ref } from "vue";

export default {
  setup() {
    const pageLoading = ref(true);

    const route = useRoute();
    onMounted(() => {
      const page = route.params.page as string;
      loaddocRaw(page);
    });

    //请求页面数据
    const loaddocRaw = (page: string): void => {
      pageLoading.value = true; 
      read({ path: page }).then((v) => {
        if (v.code==0) {
           pageLoading.value = false;
        if (v.data == "") { 
          return;
        }
        const config = Object.assign({}, BasicConfig, {
          value: v.data,
          forceAppend: true,
        });
        const cherryInstance = new Cherry(config); 
        }
       
      });
    };

    return {
      mdHeight: document.body.clientHeight - 135,
    };
  },
};
</script>

<style>
</style>
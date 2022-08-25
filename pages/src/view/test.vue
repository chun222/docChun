<!--
 * @Date: 2022-08-10 09:26:29
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-25 14:11:50
 * @FilePath: \pages\src\view\test.vue
-->
<template>
  <div id="markdown" style="height: 900px"></div>
</template>

<script lang="ts">
import Cherry from "cherry-markdown";
import { BasicConfig } from "@/tools/cherryMarkdown";
import "cherry-markdown/dist/cherry-markdown.css";

const config = Object.assign({}, BasicConfig, { value: `
:success:

**abv**

::

` });

//cherryInstance.export("pdf");
import { onMounted } from "vue";

export default {
  setup() {
    onMounted(() => {
      const cherryInstance = new Cherry(config);
    });
  },
};
</script>

<style>
</style>
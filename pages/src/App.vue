<!--
 * @Date: 2022-08-08 14:16:39
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-12 23:37:50
 * @FilePath: \pages\src\App.vue
-->
<script lang="ts">
import Layout from "@/components/layout/index.vue"; 
import { useStore } from "@/store/index";

import { computed } from "vue";
export default {
  components: {
    Layout
  },
  setup() {
    const store = useStore();

    const themeActive = computed(() => store.theme);

    return {
      themeActive
    }
  }
}
 </script>

<template> 
    <Layout :class="themeActive"></Layout>
</template> 

<style scoped>
</style>

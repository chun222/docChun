<!--
 * @Date: 2022-08-10 14:59:21
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-15 14:23:14
 * @FilePath: \pages\src\components\layout\header.vue
-->
<template>
  <a-row type="flex">
    <a-col flex="200px">LOGO</a-col>
    <a-col flex="auto">auto</a-col>
    <a-col flex="300px">
      <a-space>
        <a-switch
          :checked="themeActive"
          @change="changeTheme"
          :checkedValue="'dark'"
          :unCheckedValue="'light'"
        >
          <template #checkedChildren><moon theme="filled" /></template>
          <template #unCheckedChildren><brightness theme="filled" /></template>
        </a-switch>

        <a-input  placeholder="搜索" readonly @click="showSearch = true">
          <template #prefix>
            <search theme="outline" size="26" />
          </template>
          <template #suffix>
           
          </template>
        </a-input>
      </a-space>
    </a-col>
  </a-row>
  <vsearch :visible="showSearch" @close="showSearch = false">
 
    
  </vsearch>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs, computed } from "vue";
import { Brightness, Moon,Search } from "@icon-park/vue-next";
import { useStore } from "@/store/index";

import vsearch    from "@/components/search.vue";
export default defineComponent({
  components: {
    Brightness,
    Moon,
    Search,
    vsearch
  },
  setup() {
    const store = useStore();

    const themeActive = computed(() => store.theme);
    const state = reactive({
      showSearch:false
    });

    const changeTheme = (v) => {
      console.log(v);

      store.changeTheme(v);
    };
    return { ...toRefs(state), changeTheme, themeActive };
  },
});
</script>

<style scoped>
.ant-input-affix-wrapper { border-radius: 20px;}
</style>
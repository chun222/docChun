<!--
 * @Date: 2022-08-10 14:59:21
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-11 09:32:54
 * @FilePath: \web\src\components\layout\header.vue
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
          <template #checkedChildren><moon theme="filled"/></template>
          <template #unCheckedChildren><brightness theme="filled" /></template>
        </a-switch>
      </a-space>
    </a-col>
  </a-row>
</template>

<script>
import { defineComponent, reactive, toRefs, computed } from "vue";
import { Brightness ,Moon} from "@icon-park/vue-next";
import { useStore } from "@/store/index";
export default defineComponent({
  components: { 
    Brightness,Moon
  },
  setup() {
    const store = useStore();

    const themeActive = computed(() => store.theme);
    const state = reactive({});

    const changeTheme = (v) => {
      console.log(v);

      store.changeTheme(v);
    };
    return { ...toRefs(state), changeTheme, themeActive };
  },
});
</script>
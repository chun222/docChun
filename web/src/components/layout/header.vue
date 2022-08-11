<!--
 * @Date: 2022-08-10 14:59:21
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-10 18:01:32
 * @FilePath: \web\src\components\layout\header.vue
-->
<template>
  <a-row type="flex">
    <a-col flex="200px">LOGO</a-col>
    <a-col flex="auto">auto</a-col>
    <a-col flex="300px">
      <a-space>
        <a-switch v-model:checked="checked3" @change="changeTheme">
          <template #checkedChildren><check-outlined /></template>
          <template #unCheckedChildren><close-outlined /></template>
        </a-switch>
      </a-space>
    </a-col>
  </a-row>
</template>

<script>
import { defineComponent, reactive, toRefs } from "vue";
import { CheckOutlined, CloseOutlined } from "@ant-design/icons-vue";
import { useStore } from '@/store/index'
export default defineComponent({
  components: {
    CheckOutlined,
    CloseOutlined,
  },
  setup() {
     const store = useStore()

    const state = reactive({
      checked1: true,
      checked2: false,
      checked3: false,
    });
    const changeTheme = (v) => {
      console.log(v);
      if (v) {
        store.changeTheme("ssss")
      }else{
 store.changeTheme("vvvv")
      }
    };
    return { ...toRefs(state), changeTheme };
  },
});
</script>
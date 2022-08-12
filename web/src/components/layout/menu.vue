<!--
 * @Date: 2022-08-10 15:35:38
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-12 14:09:37
 * @FilePath: \web\src\components\layout\menu.vue
-->
<template>
  <a-menu
    v-model:openKeys="openKeys"
    v-model:selectedKeys="selectedKeys"
    mode="inline"
    :inline-collapsed="collapsed"
  >
    <template v-for="item in list" :key="item.key">
      <template v-if="!item.children">
        <a-menu-item :key="item.key"> 
          {{ item.name }}
        </a-menu-item>
      </template>
      <template v-else>
        <sub-menu :key="item.name" :menu-info="item" />
      </template>
    </template>
  </a-menu>
</template>
<script>
import { defineComponent, ref } from "vue";
import SubMenu from "./submenu.vue"; 

import { doclist } from "@/api/module/base";
const list = ref([]);
doclist({ langname: 'zh-cn' }).then((re) => {
  if (re.code == 0) {
    list.value = re.data;
  }
}); 

export default defineComponent({
  components: {
    "sub-menu": SubMenu, 
  },
  setup() {
    const collapsed = ref(false);

    const toggleCollapsed = () => {
      collapsed.value = !collapsed.value;
    };
    return {
      list,
      collapsed,
      toggleCollapsed,
      selectedKeys: ref(["1"]),
      openKeys: ref(["2"]),
    };
  },
});
</script>
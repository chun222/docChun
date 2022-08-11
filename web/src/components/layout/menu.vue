<!--
 * @Date: 2022-08-10 15:35:38
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-10 16:23:51
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
          <template #icon>
            <PieChartOutlined />
          </template>
          {{ item.title }}
        </a-menu-item>
      </template>
      <template v-else>
        <sub-menu :key="item.key" :menu-info="item" />
      </template>
    </template>
  </a-menu>
</template>
<script>
import { defineComponent, ref } from "vue";
import SubMenu from "./submenu.vue";
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  PieChartOutlined,
  MailOutlined,
} from "@ant-design/icons-vue";

const list = [
  {
    key: "1",
    title: "Option 1",
  },
  {
    key: "2",
    title: "Navigation 2",
    children: [
      {
        key: "2.1",
        title: "Navigation 3",
        children: [{ key: "2.1.1", title: "Option 2.1.1",children: [{ key: "2.1.1", title: "Option 2.1.1" ,children: [{ key: "2.1.1", title: "Option 2.1.1" }],}], }],
      },
    ],
  },
];
export default defineComponent({
  components: {
    "sub-menu": SubMenu,
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    PieChartOutlined,
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
<!--
 * @Date: 2022-08-10 08:59:35
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-11 09:35:31
 * @FilePath: \web\src\components\layout\index.vue
-->
<template>
  <a-layout>
    <a-layout-header>
      <Header></Header>
    </a-layout-header>

    <a-layout>
      <a-layout-sider
        breakpoint="md"
        collapsed-width="0"
        @collapse="onCollapse"
        @breakpoint="onBreakpoint"
      >
        <Menu></Menu>
      </a-layout-sider>
      <a-layout-content>
        <Content></Content>
      </a-layout-content>
    </a-layout>

    <a-layout-footer>
      <Footer></Footer>
    </a-layout-footer>
  </a-layout>
</template>

<script>
import { UserOutlined } from "@ant-design/icons-vue";
import { defineComponent, ref } from "vue";
import Content from "./content.vue";
import Header from "./header.vue";
import Footer from "./footer.vue";
import Menu from "./menu.vue";
export default defineComponent({
  components: {
    Content,
    Header,
    Footer,
    Menu,
    UserOutlined,
  },
  setup() {
    const onCollapse = (collapsed, type) => {
      console.log(collapsed, type);
    };

    const onBreakpoint = (broken) => {
      console.log(broken);
    };

    return {
      onCollapse,
      onBreakpoint,
    };
  },
});
</script>
<style>
</style>


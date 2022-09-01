<!--
 * @Date: 2022-08-10 08:59:35
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-09-01 11:02:17
 * @FilePath: \pages\src\components\layout\index.vue
-->
<template>
  <a-layout>
    <a-layout-header id="header">
      <Header></Header>
    </a-layout-header>

    <a-layout>
      <a-layout-sider
        breakpoint="md"
        collapsed-width="0"
        @collapse="onCollapse"
        @breakpoint="onBreakpoint"
        v-if="isIndex"
      >
        <Menu></Menu>
      </a-layout-sider>
      <a-layout-content id="content">
        <Content></Content>
        <a-layout-footer id="footer">
          <Footer></Footer>
        </a-layout-footer>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script lang="ts">
import router from "@/route/index";
import { UserOutlined } from "@ant-design/icons-vue";
import { defineComponent, ref,computed } from "vue";
import Content from "./content.vue";
import Header from "./header.vue";
import Footer from "./footer.vue";
import Menu from "./menu.vue"; 
import { useStore } from "@/store/index";
import { getMenu } from "@/tools/common";
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

    router.isReady().then(() => {
      const store = useStore(); 
      //初始化菜单什么的
      store.InitConfig().then(() => {
        console.log("初始化执行===");
        getMenu(false);
      });
    });

    return {
      isIndex: computed(()=>router.currentRoute.value.name == "/"),
      onCollapse,
      onBreakpoint,
    };
  },
});
</script>
<style>
</style>


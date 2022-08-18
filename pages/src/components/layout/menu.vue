<!--
 * @Date: 2022-08-10 15:35:38
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-18 10:29:43
 * @FilePath: \pages\src\components\layout\menu.vue
-->
<template>
  <a-menu
    v-model:selectedKeys="selectedKeys"
    mode="inline"
    :inline-collapsed="collapsed"
  >
    <template v-for="item in list" :key="item.fullpath">
      <template v-if="!item.children">
        <a-menu-item :key="item.fullpath">
          <router-link :to="`/${encodeURIComponent(item.fullpath)}`">{{
            item.name
          }}</router-link>
        </a-menu-item>
      </template>
      <template v-else>
        <sub-menu :key="item.fullpath" :menu-info="item" />
      </template>
    </template>
  </a-menu>
</template>
<script lang="ts">
import { defineComponent, ref, watch } from "vue";
import SubMenu from "./submenu.vue";
import router from "@/route/index";

import { useStore } from "@/store/index";
import { doclist } from "@/api/module/base";

import { RouteParams, useRoute } from "vue-router";

export default defineComponent({
  components: {
    "sub-menu": SubMenu,
  },
  setup() {
    const list = ref([]);
    const route = useRoute();

    const store = useStore();

    const getMenu = (isReload: boolean) => {
      doclist({ lang: store.lang.dir, version: store.version.dir }).then(
        (re) => {
          if (re.code == 0) {
            list.value = re.data;
            //默认指向第一篇文章
            if (list.value.length > 0 && isReload) {
              router.push(`/${encodeURIComponent(list.value[0].fullpath)}`);
            } else {
            }
          }
        }
      );
    };
    watch(
      () => store.version,
      (v1, v2) => {
        getMenu(true);
      }
    );

    watch(
      () => store.lang,
      (v1, v2) => {
        getMenu(true);
      }
    );

    getMenu(false);

    // store.$subscribe((mutation, state) => {
    //     console.log("state",state);
    // })

    //监听菜单选择
    const selectedKeys = ref();
    watch(
      () => route.params.page,
      (page) => {
        selectedKeys.value = [page];
      },
      { immediate: true }
    );

    const collapsed = ref(false);

    const toggleCollapsed = () => {
      collapsed.value = !collapsed.value;
    };
    return {
      list,
      collapsed,
      toggleCollapsed,
      selectedKeys,
    };
  },
});
</script>
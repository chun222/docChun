<!--
 * @Date: 2022-08-10 15:35:38
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-17 17:27:38
 * @FilePath: \pages\src\components\layout\menu.vue
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
          <router-link :to="`/${encodeURIComponent(item.fullpath)}`">{{
            item.name
          }}</router-link>
        </a-menu-item>
      </template>
      <template v-else>
        <sub-menu :key="item.name" :menu-info="item" />
      </template>
    </template>
  </a-menu>
</template>
<script lang="ts">
import { defineComponent, ref, watch } from "vue";
import SubMenu from "./submenu.vue";

import { useStore } from "@/store/index";
import { doclist } from "@/api/module/base";

export default defineComponent({
  components: {
    "sub-menu": SubMenu,
  },
  setup() {
    const list = ref([]);

    const store = useStore();
    
    const getMenu = ()=>{ 
    doclist({ lang: store.lang.dir, version: store.version.dir }).then((re) => {
      if (re.code == 0) {
        list.value = re.data;
      }
    });
    }  
    watch(() => store.version , (v1,v2) => {
     getMenu()
    })

    watch(() => store.lang , (v1,v2) => {
    getMenu()
    })

    getMenu();

     
    // store.$subscribe((mutation, state) => {
    //     console.log("state",state);
    // })

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
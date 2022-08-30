<!--
 * @Date: 2022-08-10 15:35:38
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-30 14:29:27
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
          <router-link
            :to="{
              name: dataParams.routeName ,
              params: {
                project: dataParams.project,
                lang: dataParams.lang,
                version: dataParams.version,
                page: item.relativepath,
              },
            }"
            >{{ item.name }}</router-link
          >
        </a-menu-item>
      </template>
      <template v-else>
        <sub-menu :key="item.fullpath" :menu-info="item" :dataParams="dataParams" />
      </template>
    </template>
  </a-menu>
</template>
<script lang="ts">
import { defineComponent, ref, watch, computed,reactive } from "vue";
import SubMenu from "./submenu.vue"; 
 
import { useStore } from "@/store/index";
import { RouteParams, useRoute } from "vue-router";


import { getMenu } from "@/tools/common"; 


export default defineComponent({
  components: {
    "sub-menu": SubMenu,
  },
  setup() {
    const list = computed(() => store.menus);
    const route = useRoute();

    const store = useStore();
 

    const dataParams =  reactive({
      routeName:computed(() => route.name),
      project:computed(() =>store.project.dir),
      lang:computed(() =>store.lang.dir),
      version:computed(() =>store.version.dir),
    });
 
    //对象值是否全为空
    const isEmptyObjValue = (obj: any) => {
      const values = Object.values(obj);
      for (let index = 0; index < values.length; index++) {
        const element = values[index];
        if (element != "") {
          return false;
        }
      }
      return true;
    };

    watch(
      () => store.version,
      (v1, v2) => {
        if (v2.dir != "") {
          getMenu(true);
        }
      }
    );

    watch(
      () => store.lang,
      (v1, v2) => {
        if (v2.dir != "") {
          getMenu(true);
        }
      }
    );

    // store.$subscribe((mutation, state) => {
    //     console.log("state",state);
    // })

    //监听菜单选择
    const selectedKeys = ref();
    watch(
      () => route.params.page,
      (page) => {
        selectedKeys.value = [`md/${route.params.project}/${route.params.version}/${route.params.lang}/${page}`];
       
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
      dataParams
    };
  },
});
</script>
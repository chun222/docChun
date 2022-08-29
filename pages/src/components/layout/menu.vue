<!--
 * @Date: 2022-08-10 15:35:38
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-29 16:52:48
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
              name: '/',
              params: {
                project: Params.project,
                lang: Params.lang,
                version: Params.version,
                page: item.fullpath,
              },
            }"
            >{{ item.name }}</router-link
          >
        </a-menu-item>
      </template>
      <template v-else>
        <sub-menu :key="item.fullpath" :menu-info="item" :dataParams="Params" />
      </template>
    </template>
  </a-menu>
</template>
<script lang="ts">
import { defineComponent, ref, watch, computed } from "vue";
import SubMenu from "./submenu.vue";
import router from "@/route/index";

import { doclist } from "@/api/module/base";
import { useStore } from "@/store/index";
import { RouteParams, useRoute } from "vue-router";

export default defineComponent({
  components: {
    "sub-menu": SubMenu,
  },
  setup() {
    const list = computed(() => store.menus);
    const route = useRoute();

    const store = useStore();
    const Params = ref({
      project: "",
      lang: "",
      version: "",
      page: "",
      id: "",
    });

    router.isReady().then(() => {
      const store = useStore();
      console.log("router finish!");
      //初始化菜单什么的
      store.InitConfig().then(() => {
        getMenu(false);
      });
    });

    //需解构----
    const getMenu = (isReload: boolean) => {
      const store = useStore();
      const dataParams: any = {};
      dataParams.project = store.project.dir;
      dataParams.lang = store.lang.dir;
      dataParams.version = store.version.dir;
      Params.value = dataParams;
      doclist(dataParams).then((re) => {
        if (re.code == 0) {
          //默认指向第一篇文章
          //isEmptyObjValue(router.currentRoute.value.params)
          //  const isIndex = isEmptyObjValue(router.currentRoute.value.params);
          //是否有页面
          const isNotHaspage = router.currentRoute.value.params.page == "";

      
          store.menus = re.data;
          if (re.data.length > 0) {
            if (isNotHaspage || isReload) {
              dataParams.page = re.data[0].fullpath;
            } else {
              dataParams.page = router.currentRoute.value.params.page;
            }

            //console.log("isHaspage", isNotHaspage,dataParams);
              router.push({
                name: "/",
                params: dataParams,
              });
          }
        }
      });
    };

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
      Params,
    };
  },
});
</script>
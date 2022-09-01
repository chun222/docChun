<!--
 * @Date: 2022-08-10 14:59:21
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-09-01 09:11:25
 * @FilePath: \pages\src\components\layout\header.vue
-->
<template>
  <a-row type="flex">
    <a-col flex="200px"
      ><div class="logo"><img src="@/assets/logo.png" /> chundoc</div></a-col
    >
    <a-col flex="auto">
      <!-- 横向菜单 -->
      <Menu v-if="!isIndex" :mode="'horizontal'"></Menu>
    </a-col>
    <a-col flex="600px">
      <a-space :size="20" style="flex-direction: row-reverse; float: right">
        <a-dropdown
          v-if="!isIndex"
          :getPopupContainer="(triggerNode:any) => triggerNode.parentNode"
        >
          <div class="toprightNavHover" @click.prevent>
            <config theme="outline" class="navicon" /><span>系统</span
            ><down-one theme="filled" class="navicon" />
          </div>
          <template #overlay>
            <a-menu @click="configChange">
              <a-menu-item :key="1"> 系统配置 </a-menu-item>
              <a-menu-item :key="2"> 目录配置</a-menu-item>  
            </a-menu>
          </template>
        </a-dropdown>

        <a-input
          placeholder="搜索"
          readonly
          @click="doshowSearch"
          v-if="isIndex"
        >
          <template #prefix>
            <search theme="outline" size="22" style="margin-top: 6px" />
          </template>
        </a-input>

        <a-switch
          :checked="themeActive"
          @change="changeTheme"
          :checkedValue="'dark'"
          :unCheckedValue="'light'"
        >
          <template #checkedChildren><moon theme="filled" /></template>
          <template #unCheckedChildren><brightness theme="filled" /></template>
        </a-switch>

        <!-- 语言 -->
        <a-dropdown
          :getPopupContainer="(triggerNode:any) => triggerNode.parentNode"
        >
          <div class="toprightNavHover" @click.prevent>
            <translate theme="filled" class="navicon" /><span>{{
              lang.name
            }}</span
            ><down-one theme="filled" class="navicon" />
          </div>
          <template #overlay>
            <a-menu @click="changeLang">
              <a-menu-item v-for="v in langs" :key="v">{{
                v.name
              }}</a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>

        <!-- 版本 -->
        <a-dropdown
          :getPopupContainer="(triggerNode:any) => triggerNode.parentNode"
        >
          <div class="toprightNavHover" @click.prevent>
            <span>{{ version.name }}</span
            ><down-one theme="filled" class="navicon" />
          </div>
          <template #overlay>
            <a-menu @click="changeVersion">
              <a-menu-item v-for="v in versions" :key="v">{{
                v.name
              }}</a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>

        <!-- 项目 -->
        <a-dropdown
          :getPopupContainer="(triggerNode:any) => triggerNode.parentNode"
        >
          <div class="toprightNavHover" @click.prevent>
            <span>{{ project.name }}</span
            ><down-one theme="filled" class="navicon" />
          </div>
          <template #overlay>
            <a-menu @click="changeProject">
              <a-menu-item v-for="v in projects" :key="v">{{
                v.name
              }}</a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </a-space>
    </a-col>
  </a-row>
  <vsearch
    :visible="showSearch"
    :parentNode="domParent"
    @close="showSearch = false"
  >
  </vsearch>
 

  <ConfigDrawer :visible="showConfig" @close="showConfig=false"></ConfigDrawer>
</template>

<script lang="ts">
import router from "@/route/index";
import { defineComponent, reactive, toRefs, computed, nextTick } from "vue";
import ConfigDrawer from "./drawer/config.vue";
import {
  Brightness,
  Moon,
  Search,
  DownOne,
  Translate,
  Config,
} from "@icon-park/vue-next";
import { useStore } from "@/store/index";
import vsearch from "@/components/search.vue";
import Menu from "./menu.vue";
export default defineComponent({
  components: {
    ConfigDrawer,
    Brightness,
    Moon,
    Search,
    vsearch,
    DownOne,
    Translate,
    Config,
    Menu,
  },
  setup() {
 
    const store = useStore();

    const themeActive = computed(() => store.theme);
    const state = reactive({
      showSearch: false,
      domParent: null,
      showConfig:false
    });

    const doshowSearch = () => {
      nextTick(() => {
        state.domParent = document.getElementById("content");
        state.showSearch = true;
      });
    };
    const configChange = (v: any) => {
      console.log(v); 
      //系统配置
      if (v.key==1) {
          state.showConfig = true
      }else if(v.key==2){
        //目录配置
        
      }
    };

    const changeTheme = (v: string) => { 
      store.changeTheme(v);
    };

    const changeLang = ({ key }) => {
      store.changeLang(key);
    };

    const changeVersion = ({ key }) => {
      store.changeVersion(key);
    };

    const changeProject = ({ key }) => {
      store.changeProject(key);
    };

    const langs = computed(() => store.langs);
    const projects = computed(() => store.projects);
    const versions = computed(() => store.versions);
    const lang = computed(() => store.lang);
    const version = computed(() => store.version);
    const project = computed(() => store.project);
    return {
      ...toRefs(state),
      changeTheme,
      themeActive,
      langs,
      versions,
      lang,
      version,
      projects,
      project,
      changeLang,
      changeVersion,
      changeProject,
      doshowSearch,
      configChange,
      isIndex: computed(() => router.currentRoute.value.name == "/"),
    };
  },
});
</script>

 
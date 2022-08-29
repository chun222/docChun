<!--
 * @Date: 2022-08-10 14:59:21
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-29 10:37:02
 * @FilePath: \pages\src\components\layout\header.vue
-->
<template>
  <a-row type="flex">
    <a-col flex="200px"><div class="logo"><img src="@/assets/logo.png" />  chundoc</div></a-col>
    <a-col flex="auto">auto</a-col>
    <a-col flex="600px">
      <a-space :size="20" style="flex-direction: row-reverse; float: right">
        <a-input placeholder="搜索" readonly @click="doshowSearch">
          <template #prefix>
            <search theme="outline" size="22" style="margin-top:6px"/>
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
      </a-space>
    </a-col>
  </a-row>
  <vsearch :visible="showSearch"  :parentNode="domParent" @close="showSearch = false"> </vsearch>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs, computed,nextTick } from "vue";
import {
  Brightness,
  Moon,
  Search,
  DownOne,
  Translate,
} from "@icon-park/vue-next";
import { useStore } from "@/store/index";
import vsearch from "@/components/search.vue";  
export default defineComponent({
  components: {
    Brightness,
    Moon,
    Search,
    vsearch,
    DownOne,
    Translate,
  },
  setup() {

    
   
    const store = useStore();

    const themeActive = computed(() => store.theme);
    const state = reactive({
      showSearch: false,
      domParent:null
    });

     const doshowSearch = ()=>{
      nextTick(()=>{ 
         state.domParent = document.getElementById('content') ;
          state.showSearch = true ;
      })
    }

    const changeTheme = (v: string) => {
      console.log(v);
      store.changeTheme(v);
    };

    const changeLang= ({ key }) => {
      store.changeLang(key)
    };

    const changeVersion = ({ key }) => {
     store.changeVersion(key,true)
    };

    const langs = computed(() => store.langs);

    const versions = computed(() => store.versions);
    const lang = computed(() => store.lang);
    const version = computed(() => store.version); 
    return {
      ...toRefs(state),
      changeTheme,
      themeActive,
      langs,
      versions,
      lang,
      version,
      changeLang,
      changeVersion, 
      doshowSearch,
    };
  },
});
</script>

 
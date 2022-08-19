<!--
 * @Date: 2021-12-17 09:44:58
 * @LastEditors: 春贰
 * @Desc: 
 * @LastEditTime: 2022-08-19 11:20:51
 * @FilePath: \pages\src\components\search.vue
-->
<template>
  <a-modal
    :getContainer="parentNode"
    :footer="null"
    :visible="visible"
    :cancelText="'cancel'"
    :okText="'submit'"
    @cancel="cancel()"
    :closable="false"
    :destroyOnClose="true"
    :width="600"
  >
    <template #title>
      <a-input
        ref="searchRef"
        id="searchRef"
        @change="changeInput"
        v-model:value="state.serchValue"
        placeholder="搜索文档"
      >
        <template #prefix>
          <search theme="outline" size="22" style="margin-top: 6px" />
        </template>
      </a-input>
    </template>
    <div>
      <a-empty v-if="state.loadingType == 0" :image="simpleImage" />
      <div v-else-if="state.loadingType == 1" style="text-align: center">
        <a-spin size="large" />
      </div>
      <div v-else class="search_result">
        <ul>
          <router-link
            v-for="item in state.searchResult"
            :to="`/${encodeURIComponent(item.PagePath)}/${encodeURIComponent(
              item.Id
            )}`"
            @click="cancel()"
          >
            <li>
              <h2><view-list class="mt5" theme="outline" size="16"/>&nbsp;{{ item.PageName }} <span><pound theme="outline" size="14" /><span v-html="item.showId"></span></span></h2> 
              <p v-html="item.Text"></p>
            </li>
          </router-link>
        </ul>
      </div>
    </div>
  </a-modal>
</template>
<script lang="ts">
import {
  defineComponent,
  onMounted,
  reactive,
  ref,
  toRaw,
  watch,
  onUpdated,
} from "vue";
import { debounce } from "@/tools/common";
import { useStore } from "@/store/index";
import { search } from "@/api/module/base";
import { Search,ViewList,Pound } from "@icon-park/vue-next";
import { Empty } from "ant-design-vue";
export default defineComponent({
  props: {
    visible: {
      type: Boolean,
    }, 
    parentNode:{
      type:HTMLElement
    }
  },
  components: {
    Search,
    ViewList,
    Pound
  },
  emit: ["close"],
  setup(props, context) {
    console.log(props.parentNode);
    const searchRef = ref<HTMLElement>();
    const store = useStore();
    const state = reactive({
      depts: [],
      showPath: true,
      treeselect: [],
      serchValue: "",
      searchResult: [],
      loadingType: 0,
    });

    //获取不到a-modal里的input是否显示 不得已setTimeout这么做
    watch(
      () => props.visible,
      (v) => {
        if (v) {
          state.searchResult = [];
          state.loadingType = 0;
          setTimeout(() => {
            if (document.getElementById("searchRef")) {
              document.getElementById("searchRef").focus();
            }
          }, 100);
        }
      }
    );

    // onUpdated(() => {
    //   console.log("onUpdated");
    //   console.log(searchRef.value);
    // })

    const cancel = (e = false) => {
      state.serchValue = "";
      context.emit("close", e);
    };
    const changeInput = debounce(() => {
      if (state.serchValue.trim()=="") {
        state.loadingType = 0;
        return 
      }
      state.loadingType = 1;
      search({
        lang: store.lang.dir,
        version: store.version.dir,
        keyword: state.serchValue,
      }).then((v) => { 
        if (v.data) {
           state.loadingType = 2;
           
            const rule= new RegExp(state.serchValue)
           v.data.forEach((item:any,i:number) => {
             v.data[i]["showId"] = item.Id.replace(rule,`<span class="searched">${state.serchValue}</span>`)
             v.data[i]["Text"] = item.Text.replace(rule,`<span class="searched">${state.serchValue}</span>`)
           });

           console.log(v.data,"v.data");
           state.searchResult = v.data;
        } else {
          state.loadingType = 0;
           state.searchResult = [];
        }
       
      });
    }, 500);

    return {
      changeInput,
      cancel,
      state,
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
      searchRef,
      simpleImage: Empty.PRESENTED_IMAGE_SIMPLE,
    };
  },
});
</script>
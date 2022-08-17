<!--
 * @Date: 2021-12-17 09:44:58
 * @LastEditors: 春贰
 * @Desc: 
 * @LastEditTime: 2022-08-17 17:33:52
 * @FilePath: \pages\src\components\search.vue
-->
<template>
  <a-modal
    :footer="null"
    :visible="visible"
    :cancelText="'cancel'"
    :okText="'submit'"
    @cancel="cancel()"
    :closable="false"
  >
    <template #title>
      <a-input @change="changeInput" v-model:value="state.serchValue"></a-input>
    </template>
    <div>
      <ul>
        <router-link
          v-for="item in state.searchResult"
          :to="`/${encodeURIComponent(item.PagePath)}/${encodeURIComponent(
            item.Id
          )}`"
          
          @click="cancel()"
        >
          <li>
            <h2>{{ item.PageName }}</h2>
            <h5>{{ item.Id }}</h5>
            <p>{{ item.Text }}</p>
          </li>
        </router-link>
      </ul>
    </div>
  </a-modal>
</template>
<script lang="ts"> 
import { defineComponent, reactive, ref, toRaw, watch } from "vue";
import { debounce } from "@/tools/common"; 
import { useStore } from "@/store/index";
import { search } from "@/api/module/base";
export default defineComponent({
  props: {
    visible: {
      type: Boolean,
    },
    treeselect: {
      type: Array,
    },
    parent: {
      type: Number,
    },
  },
  emit: ["close"],
  setup(props, context) {
     const store = useStore();
    const state = reactive({
      depts: [],
      showPath: true,
      treeselect: [],
      serchValue: "",
      searchResult: [],
    });

    const cancel = (e = false) => {
      state.serchValue = "";
      context.emit("close", e);
    };
    const changeInput = debounce(() => {
      search({ lang: store.lang.dir, version:store.version.dir, keyword: state.serchValue }).then((v) => { 
        state.searchResult = v.data;
      });
    }, 500);

    return {
      changeInput,
      cancel,
      state,
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
    };
  },
});
</script>
<!--
 * @Date: 2022-08-10 09:26:29
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-09-01 11:31:15
 * @FilePath: \pages\src\view\admin\admin.vue
-->
<template>

<div>
  <div id="markdown" :style="{ height: mdHeight + 'px' }"></div>

  <!-- 曲线救国 标准方法获取不到 -->
  <div class="diyButtons">
    <span
      class="diyButtonspan cherry-toolbar-dropdown cherry-toolbar-button"
      @click="viewpage"
      >查看</span
    >
    <span
      class="diyButtonspan cherry-toolbar-dropdown cherry-toolbar-button"
      @click="saveMd(true)"
      >保存</span
    >
  </div>
</div>
  
</template>

<script lang="ts">
import Cherry from "cherry-markdown";
import { CherryMarkdownConfig } from "@/tools/cherryMarkdown";
import "cherry-markdown/dist/cherry-markdown.css";

import { RouteParams, useRoute } from "vue-router";
import { useStore } from "@/store/index";

import { read, savecontent } from "@/api/module/base";
import {TypePath} from '@/model';
//cherryInstance.export("pdf");
import {
  onMounted,
  ref,
  watch,
  reactive,
  getCurrentInstance,
  onUnmounted,
  defineComponent,
} from "vue";
import "font-awesome/css/font-awesome.min.css";

import router from "@/route/index";
 

export default {
  components:{
     
  },
  setup() { 
    const state = reactive({ 
    })
    const globleConfig = getCurrentInstance()?.proxy;
    const pageLoading = ref(true);

    const store = useStore();
    const route = useRoute();
    const BConfigClass = new CherryMarkdownConfig({
      forceAppend: true,
    }).BasicConfig;

    let cherryInstance: any;

    let IntervalKey: NodeJS.Timer;

    //当前页面
    let curpage: TypePath = {
      page: "",
      project: store.project.dir,
      version: store.version.dir,
      lang: store.lang.dir,
    };

    watch(
      () => route.params,
      (val: any, val2: any) => {
        console.log("改变了参数",val,val2);
        loaddocRaw(val.page as string);
      }
    );

    onMounted(() => {
      cherryInstance = new Cherry(BConfigClass);

      document
        .querySelector(".cherry-toolbar")
        .appendChild(document.querySelector(".diyButtons"));
      // document.querySelector(".diyButtons").childNodes.forEach((v)=>{
      //   document.querySelector(".cherry-toolbar").appendChild(v)
      // })
      // document.querySelector(".diyButtons").remove()

      //自动保存
      IntervalKey = setInterval(() => {
        // saveMd(false)
      }, 1000);
    });

    onUnmounted(() => {
      clearInterval(IntervalKey);
    });

    watch(
      () => store.InitOk,
      (v1: boolean, v2: boolean) => {
        const page = route.params.page as string;
        loaddocRaw(page);
      }
    );

    //请求页面数据
    const loaddocRaw = (page: string): void => {
      pageLoading.value = true;
      curpage = {
        page: page,
        project: store.project.dir,
        version: store.version.dir,
        lang: store.lang.dir,
      };
      read(curpage).then((v) => {
        if (v.code == 0) {
          pageLoading.value = false;
          cherryInstance.setMarkdown(v.data, false);
        }
      });
    };

    const viewpage = () => {
      const urlpage = router.resolve({
        name: "/",
        params: curpage,
      });
      window.open(urlpage.href, "_blank");
    };

    const saveMd = (showmsg: boolean) => {
      curpage.content = cherryInstance.getMarkdown();
      // console.log(cherryInstance.getMarkdown());
      savecontent(curpage).then((re) => {
        if (re.code == 0) {
          if (showmsg) {
            globleConfig.$notice.success({
              message: "保存成功~",
            });
          }
        } else {
          globleConfig.$notice.error({
            message: re.msg,
          });
        }
      });
    };

    return {
      state,
      mdHeight: document.body.clientHeight - 135,
      saveMd,
      viewpage,
    };
  },
};
</script>

<style>
</style>
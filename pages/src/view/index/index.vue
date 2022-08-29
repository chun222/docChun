<!--
 * @Date: 2022-08-08 14:16:39
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-29 17:47:50
 * @FilePath: \pages\src\view\index\index.vue
-->
<template>
  <div class="cherry-previewer doc-previewer cherry-markdown">
    <template v-if="!pageLoading">
      <div class="doc-card" v-html="rawHtml" ref="doc_card"></div>

      <div class="title_nav">
        <p v-for="v in listTitle">
          <router-link
            :to="{
              name: '/',
              params:v.params,
            }"
            title="heading"
            :style="{ 'margin-left': v.level + 'px' }"
            >{{ v.nodeName }}</router-link
          >
        </p>
      </div>
    </template>
    <div class="content-loading" v-else>
      <a-spin size="large" />
    </div>

    <div @click="jumpLocation('', 0)" v-show="showTop" class="totop">
      <to-top theme="outline" size="34" />
    </div>
  </div>
</template>

<script setup lang="ts">
 
import { svgArr } from "./svg";
import { ref, onMounted, nextTick, watch, defineComponent } from "vue"; 
import DOMPurify from "dompurify"; 
import { RouteParams, useRoute } from "vue-router";
import { read } from "@/api/module/base";
import { useStore } from "@/store/index";
import { ToTop } from "@icon-park/vue-next"; 
 
import Cherry from "cherry-markdown";
import "cherry-markdown/dist/cherry-markdown.css";

import { BasicConfig } from "@/tools/cherryMarkdown";
const route = useRoute();
const store = useStore();
watch(
  () => route.params.page,
  (val: any) => {
    loaddocRaw(val as string);
  }
);

watch(
  () => route.params.id,
  (val: any) => {
    jumpLocation(val as string);
  }
);

defineComponent({
  ToTop,
});

defineProps({
  msg: String,
});

const pageLoading = ref(true);

const count = ref(0);
const doc_card = ref();
const listTitle = ref([]);
 
 
//跳转
const jumpLocation = (id: string, top?: number) => {
  if (top != undefined) {
    document.getElementById("content").scroll({ top: top, behavior: "smooth" });
    return;
  }
  if (document.getElementById(id)) {
    const el = document.getElementById(id);
    document.getElementById("content").scroll({
      top: top == undefined ? el.offsetTop : top,
      behavior: "smooth",
    });
    // document.getElementById(id).scrollIntoView({
    //   behavior: "smooth",
    // });
  }
};

  
const rawHtml = ref("");
//请求页面数据
const loaddocRaw = (page: string): void => {
  pageLoading.value = true;
  listTitle.value = [];
  read({ path: page }).then((v) => {
    pageLoading.value = false;
    if (v.data == "") {
      rawHtml.value = "404";
      return;
    }
    const config = Object.assign({}, BasicConfig,{value:v.data, forceAppend: true}) 
  const cherryInstance = new Cherry(config);
     rawHtml.value = cherryInstance.getHtml()
 
    nextTick(() => {
      document.getElementById("markdown").remove()

      console.log(cherryInstance.getToc());
      //原生
      // const Hdomlist = doc_card.value.querySelectorAll(
      //   "h1, h2, h3, h4, h5, h6"
      // );

      //自带api
      const Hdomlist  = cherryInstance.getToc()
      Hdomlist.forEach((x) => {
        listTitle.value.push({
          page: page,
          id: x.id,
          nodeName: decodeURI(x.id),
          //parseInt(x.nodeName.replace(/[Hh]/, ""))  //自带需处理
          level: x.level * 10,
          params:Object.assign({},route.params,{id:x.id})    
        });
      });

      if (route.params.id) {
        jumpLocation(route.params.id as string);
      }

      // let clipboard = null;
      // //处理代码复制
      // document.querySelectorAll("code .hljs-box").forEach((element) => {
      //   const butDom = element.querySelector(".doc-codelang");
      //   const copyTextS = decodeURIComponent(butDom.getAttribute("value")); 
      //   butDom.addEventListener("click", () => {
      //     copyText(copyTextS)
      //     butDom.innerHTML = "复制成功";
      //     // clipboard = new Clipboard(butDom, {
      //     //   text: function (trigger) {
      //     //     return copyText;
      //     //   },
      //     // });
      //     // clipboard.on("success", function (e) {
      //     //   butDom.innerHTML = "复制成功";
      //     //   e.clearSelection();
      //     // });
      //   });
      //   const oldcodelang = butDom.innerHTML;
      //   element.addEventListener("mouseover", () => {
      //     butDom.innerHTML = "复制";
      //   });
      //   element.addEventListener("mouseout", () => {
      //     butDom.innerHTML = oldcodelang;
      //   });
      // });
    });
  });
};

const showTop = ref(false);

const lisscroll = () => {
  const odiv = document.getElementById("content");
  odiv.onscroll = function (e) {
    if (odiv.scrollTop > 200) {
      showTop.value = true;
    } else {
      showTop.value = false;
    }
  };
};

onMounted(() => {
  lisscroll();
  const page = route.params.page as string;
  loaddocRaw(page);
});
</script>

 
<style scoped>
</style>

<!--
 * @Date: 2022-08-08 14:16:39
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-18 16:14:10
 * @FilePath: \pages\src\view\index\index.vue
-->
<template>
  <div class="doc-main" id="doc-main">
    <template v-if="!pageLoading">
      <div class="doc-card" v-html="rawHtml" ref="doc_card"></div>

      <div class="title_nav">
        <p v-for="v in listTitle">
          <router-link
            :to="`/${encodeURIComponent(v.page)}/${encodeURIComponent(v.id)}`"
            title="heading"
            :style="{ 'margin-left': v.level + 'px' }"
            >{{ v.id }}</router-link
          >
        </p>
      </div>
    </template>
    <div class="content-loading" v-else>
      <a-spin size="large" />
    </div>

    <div @click="jumpLocation('doc-main')" v-show="showTop" class="totop"><to-top theme="outline" size="34" /></div>
  </div>
</template>

<script setup lang="ts">
import { svgArr } from "./svg";
import { ref, onMounted, nextTick, watch, defineComponent } from "vue";
import { marked } from "marked";
import DOMPurify from "dompurify";
import hljs from "highlight.js";
import "highlight.js/styles/github.css";
import "highlight.js/styles/github-dark.css";
import { RouteParams, useRoute } from "vue-router";
import { read } from "@/api/module/base";
import { useStore } from "@/store/index";
import { ToTop } from "@icon-park/vue-next";
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
const jumpLocation = (id: string) => {
  if (document.getElementById(id)) {
    document.getElementById(id).scrollIntoView({
      behavior: "smooth",
    });
  }
};

// Set options
// `highlight` example uses https://highlightjs.org
marked.setOptions({
  baseUrl: "https://marked.js.org/aaa/",
  renderer: new marked.Renderer(),
  highlight: function (code, lang) {
    const language = hljs.getLanguage(lang) ? lang : "plaintext";
    return hljs.highlight(code, { language }).value;
  },
  langPrefix: "hljs language-", // highlight.js css expects a top-level 'hljs' class.
  pedantic: false,
  gfm: true,
  breaks: false,
  sanitize: false,
  smartLists: true,
  smartypants: false,
  xhtml: false,
});

const tips = {
  name: "tips",
  level: "block", // Is this a block-level or inline-level tokenizer?
  start(src) {
    // console.log( src.match(/:tip:\n/)?.index,"start",src);
    return src.match(/:(info|success|warning|error):/)?.index;
  }, // Hint to Marked.js to stop and check for a match /:tip:[^:\n]/
  tokenizer(src, tokens) {
    const rule =
      /^(?::(info|success|warning|error):[\n\s])[\S\s]*?(?<=::[\n\s])/; //  /^(?::tip:[\n\s])[\S\s]*?(?<=::[\n\s])/
    const match = rule.exec(src);

    if (match) {
      const token = {
        // Token to generate
        type: "tips", // Should match "name" above
        raw: match[0], // Text to consume from the source
        text: match[0].trim(), // Additional custom properties
        class: match[1],
        tokens: [], // Array where child inline tokens will be generated
      };
      //  this.lexer.inlineTokens(token.text, token.tokens);
      this.lexer.inlineTokens(token.text, token.tokens); // Queue this data to be processed for inline tokens  blockTokens
      return token;
    }
  },
  renderer(token) {
    // console.log(this.parser.parseInline(token.tokens).replace(/^:(info|success|warning|error):/,"").replace(/::$/,""));
    // console.log("redendr",this.parser.parseInline(token.tokens)); .replace(/^:tip:/,"").replace(/::$/,"")
    return `<div class="${token.class}" >
    <div class="tip-header">${svgArr[token.class] || ""} <span>${
      token.class
    }</span></div>
    <p> ${this.parser
      .parseInline(token.tokens)
      .replace(/^:(info|success|warning|error):\s+/, "")
      .replace(/::$/, "")}</p>
    </div>`; // parseInline to turn child tokens into HTML   (\n|\r|\r\n|↵)/g
  },
};

marked.use({ extensions: [tips] });
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
    rawHtml.value = DOMPurify.sanitize(marked.parse(v.data));
    nextTick(() => {
      (window as any).MathJax.typeset();
      const domlist = doc_card.value.querySelectorAll("h1, h2, h3, h4, h5, h6");
      domlist.forEach((x) => {
        listTitle.value.push({
          page: page,
          id: x.id,
          nodeName: x.nodeName,
          level: parseInt(x.nodeName.replace(/[Hh]/, "")) * 10,
        });
      });

      if (route.params.id) {
        jumpLocation(route.params.id as string);
      }
    });
  });
};

const showTop = ref(false)

const lisscroll = () => { 
    const odiv = document.getElementById("content"); 
    odiv.onscroll = function (e) {
       if (odiv.scrollTop>200) {
          showTop.value = true
       } else {
         showTop.value = false
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

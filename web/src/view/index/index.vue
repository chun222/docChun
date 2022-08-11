<!--
 * @Date: 2022-08-08 14:16:39
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-11 15:12:41
 * @FilePath: \web\src\view\index\index.vue
-->
<template>  
    <a-row type="flex" :gutter="10"> 
    <a-col flex="auto">
         <div class="doc-card" v-html="rawHtml" ref="doc_card"></div>
    </a-col>
    <a-col flex="300px">
        <div class="title_nav">  
        <p v-for="v in listTitle"><router-link  :to="'/'+v.id" title="heading" :style="{'margin-left':v.level+'px'}">{{v.id}}</router-link></p> 
        </div></a-col>
  </a-row>
 
</template>

<script setup>
import { useStore } from '../../store/index' 
import {svgArr} from './svg'
import { ref,onMounted,nextTick,watch } from "vue";
import { marked } from "marked";
import DOMPurify from "dompurify";
import hljs from "highlight.js";
import "highlight.js/styles/github.css";
import "highlight.js/styles/github-dark.css"; 
import { useRoute } from 'vue-router';
 
const route = useRoute();  
console.log(route.params);
 
watch(
    () => route.params,
    (val) => {
       jumpLocation(val.id)
    },
);
defineProps({
  msg: String,
});

const count = ref(0);
const doc_card = ref() 
const listTitle = ref([])  

//跳转
const jumpLocation = (id)=>{ 
  if (document.getElementById(id)) {
     document.getElementById(id).scrollIntoView({
behavior: "smooth", 
})
  }
 
}

 onMounted(() => { 
		nextTick(() => { 
      MathJax.typeset(); 
      const domlist = doc_card.value.querySelectorAll("h1, h2, h3, h4, h5, h6") 
       domlist.forEach((x)=>{
      listTitle.value.push( {id:x.id,nodeName:x.nodeName,level:parseInt(x.nodeName.replace(/[Hh]/,""))*10 })
     }) 
		})
 
    });

// Set options
// `highlight` example uses https://highlightjs.org
marked.setOptions({
  baseUrl:"https://marked.js.org/aaa/",
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
  name: 'tips',
  level: 'block',                                     // Is this a block-level or inline-level tokenizer?
  start(src) { 
   // console.log( src.match(/:tip:\n/)?.index,"start",src);
    return src.match(/:(info|success|warning|error):/)?.index;
     }, // Hint to Marked.js to stop and check for a match /:tip:[^:\n]/
  tokenizer(src, tokens) {
    const rule = /^(?::(info|success|warning|error):[\n\s])[\S\s]*?(?<=::[\n\s])/;    //  /^(?::tip:[\n\s])[\S\s]*?(?<=::[\n\s])/
    const match = rule.exec(src);

    if (match) { 
      const token = {                                 // Token to generate
        type: 'tips',                      // Should match "name" above
        raw: match[0],                                // Text to consume from the source
        text: match[0].trim(),                        // Additional custom properties
        class: match[1], 
        tokens: []                                    // Array where child inline tokens will be generated
      };
     //  this.lexer.inlineTokens(token.text, token.tokens); 
      this.lexer.inlineTokens(token.text, token.tokens);    // Queue this data to be processed for inline tokens  blockTokens
      return token;
    }
  },
  renderer(token) {
   // console.log(this.parser.parseInline(token.tokens).replace(/^:(info|success|warning|error):/,"").replace(/::$/,""));
   // console.log("redendr",this.parser.parseInline(token.tokens)); .replace(/^:tip:/,"").replace(/::$/,"")
    return `<div class="${token.class}" >
    <div class="tip-header">${svgArr[token.class]||""} <span>${token.class}</span></div>
    <p> ${this.parser.parseInline(token.tokens).replace(/^:(info|success|warning|error):\s+/,"").replace(/::$/,"")}</p>
    </div>`; // parseInline to turn child tokens into HTML   (\n|\r|\r\n|↵)/g
  }
};
 
marked.use({ extensions: [tips] }); 

const rawHtml = marked.parse(`
:info:

info 
**ssss**

ss
:: 
:success: 
success 
::
:warning: 
warning 
::
:error: 
error 
::  
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题 

*斜体文本*
_斜体文本_
**粗体文本**
__粗体文本__
***粗斜体文本***
___粗斜体文本___ 

***

* * *

*****

- - - 

----------

~~删除线~~

<u>下划线</u>

* 第一行
* 第二行

- 第一行
- 第二行

+ 第一行
+ 第二行

1.第一行
  - 第一点
  - 第二点

 

> 一级引用
> > 二级引用
> > > 三级引用

> 区块中使用列表
> 1. 第一项
> 2. 第二项
>> + 第一项
>> + 第二项
>> + 第三项

[这是百度](https://www.baidu.com)

 ###### 哈哈哈 

\`\`\`js
   protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
    }
\`\`\` 

> #### 这是一个四级标题 


> 1. 这是第一行列表项
> 2. 这是第二行列表项

:success:

abv

::

:error:

**ssss**

### 事实上

Topic 1:  D
escription 1   
:Topic 1:  Description 1   

::   
|标题|标题|标题|
|:---|:---:|---:|
|居左测试文本|居中测试文本|居右测试文本|
|居左测试文本1|居中测试文本2|居右测试文本3|
|居左测试文本11|居中测试文本22|居右测试文本33|
|居左测试文本111|居中测试文本222|居右测试文本333|

![中文](./docimg/5.png)

![中文](http://localhost:8082/src/assets/vue.svg)
 

## 混合 
:warning: 
Topic 1:  Description 1   
:Topic 1:  Description 1  
::
$$x = {-b \pm \sqrt{b^2-4ac} \over 2a}.$$
# 事实上
啊`);


 
</script>

 
<style scoped>

</style>

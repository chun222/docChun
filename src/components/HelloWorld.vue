<!--
 * @Date: 2022-08-08 14:16:39
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-09 17:38:12
 * @FilePath: \web\src\components\HelloWorld.vue
-->
<script setup>
import { ref,onMounted,nextTick } from "vue";
import { marked } from "marked";
import DOMPurify from "dompurify";
import hljs from "highlight.js";
import "highlight.js/styles/github.css";
import "highlight.js/styles/github-dark.css";
defineProps({
  msg: String,
});

const count = ref(0);



 onMounted(() => { 
		nextTick(() => { 
      MathJax.typeset();

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
      this.lexer.inlineTokens(token.text, token.tokens); 
      //this.lexer.inline(token.text, token.tokens);    // Queue this data to be processed for inline tokens  blockTokens
      return token;
    }
  },
  renderer(token) {
    console.log(token);
   // console.log("redendr",this.parser.parseInline(token.tokens)); .replace(/^:tip:/,"").replace(/::$/,"")
    return `<div class="${token.class}" >${this.parser.parseInline(token.tokens).replace(/^:(info|success|warning|error):/,"").replace(/::$/,"")}</div>`; // parseInline to turn child tokens into HTML 
  }
};
 
marked.use({ extensions: [tips] }); 

const rawHtml = marked.parse(`
水水水水

> 这是一段引用    //在\`>\`后面有 1 个空格
> 
>     这是引用的代码块形式    
>     
> 代码例子：

\`\`\`js
   protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
    }
\`\`\`
  
 

> 一级引用
> > 二级引用
> > > 三级引用

> #### 这是一个四级标题
> 
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



## 混合

:warning: 
Topic 1:  Description 1   
:Topic 1:  Description 1  
::
$$x = {-b \pm \sqrt{b^2-4ac} \over 2a}.$$
# 事实上
啊`);



// `
// # Marked in Node.js
// Rendered by **marked**.

// :tip
// sss
// :
// \`\`\`go
// 	for _, v := range pointPrices {
// 			if _, ok := pointPricesMap[uint(v.UnitId)]; !ok {
// 				pointPricesMap[uint(v.UnitId)] = make([]AppDbModel.EmsUnitPrice, 0)
// 			}
// 			pointPricesMap[uint(v.UnitId)] = append(pointPricesMap[uint(v.UnitId)], v)
// 		}
// function(){
//   alert(1)
// }
// \`\`\`
// `
</script>

<template>
  <h1>{{ msg }}</h1>

  <div class="card" v-html="rawHtml"></div>
</template>

<style>
body {
    font-family: -apple-system,BlinkMacSystemFont,"Segoe UI",Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Segoe UI Symbol";
    font-size: 16px;
    line-height: 1.5;
    word-wrap: break-word;
    background: #F9F9F9;
}

#container {
    max-width: 900px;
    margin: auto;
}

#content {
    padding: 5px 10px;
    border: 1px solid #ddd;
    border-radius: 3px;
    background: white;
}

#content h1:first-child {
    margin-top: 0px;
}

nav {
    border: 1px solid #ddd;
    border-radius: 3px;
    background: white;
    margin-right: 10px;
}

nav > ul {
    position: sticky;
    top: 5px;
    margin: 10px 0px;
    padding: 0;
    list-style-type: none;
    font-size: 14px;
}

nav > ul > li {
    min-width: 125px;
    padding: 0px 15px;
}

nav > ul > li > ul {
    padding-left: 25px;
}

nav > ul > li > ul > li {
    font-size: 0.8em;
}

nav .selected {
    color: #111;
    font-weight: bold;
}

nav .selected:hover {
    text-decoration: none;
}

header {
    display: flex;
}

header h1 { margin: 0; }

table {
    border-spacing: 0;
    border-collapse: collapse;
    border: 1px solid #ddd;
}

td, th {
    border: 1px solid #ddd;
    padding: 5px;
}

a {
    color: #0366d6;
    text-decoration: none;
}

a:hover {
    text-decoration: underline;
}

pre {
    font-family: "SFMono-Regular",Consolas,"Liberation Mono",Menlo,Courier,monospace;
    padding: 16px;
    overflow: auto;
    font-size: 85%;
    line-height: 1.45;
    background-color: #f6f8fa;
    border-radius: 3px;
}

code:not([class]) {
    padding: 0.2em 0.4em;
    margin: 0;
    font-size: 85%;
    background-color: rgba(27,31,35,0.05);
    border-radius: 3px;
}

.github-corner:hover .octo-arm{animation:octocat-wave 560ms ease-in-out}@keyframes octocat-wave{0%,100%{transform:rotate(0)}20%,60%{transform:rotate(-25deg)}40%,80%{transform:rotate(10deg)}}@media (max-width:500px){.github-corner:hover .octo-arm{animation:none}.github-corner .octo-arm{animation:octocat-wave 560ms ease-in-out}}
.error { color: red;}
.success { color: aquamarine;}
.warning { color: rgb(212, 200, 19);}
.read-the-docs {
  color: #888;
}
</style>

<!--
 * @Date: 2022-08-08 14:16:39
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-09 14:41:54
 * @FilePath: \web\src\components\HelloWorld.vue
-->
<script setup>
import { ref } from "vue";
import { marked } from "marked";
import DOMPurify from "dompurify";
import hljs from "highlight.js";
import "highlight.js/styles/github.css";
import "highlight.js/styles/github-dark.css";
defineProps({
  msg: String,
});

const count = ref(0);

// Set options
// `highlight` example uses https://highlightjs.org
marked.setOptions({
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
console.log(`::tip
 
abcdefgsssaaa
::`.replace(/::tip/,"").replace(/::/,""));
// 使用正则
// let reg = /(:tip)/;
// let str = `:tip
// sss
// :`;
// // 使用
// console.log(str.match(reg));
// str.match(reg); // 输出 ['1234']

// const descriptionList = {
//   name: "descriptionList",
//   level: "block", // Is this a block-level or inline-level tokenizer?
//   start(src) {  
//     return src.match(/:[^:\n]/)?.index;
//   }, // Hint to Marked.js to stop and check for a match
//   tokenizer(src, tokens) {
//     // console.log(src,"===tokenizersrc");
//     const rule = /^(?::tip:[^:\n]*(?:\n|$))+/; // /^(?:::tip[\n\s])[\s\S]*(?=::\n)/
//     const match = rule.exec(src);
//     // const match = src.match(rule);
   
//     if (match) {
//        console.log(src, match, "match");
//       const token = {
//         // Token to generate
//         type: "descriptionList", // Should match "name" above
//         raw: match[0], // Text to consume from the source
//         text: match[0].trim(), // Additional custom properties .replace(/::tip/,"")
//         tokens: [], // Array where child inline tokens will be generated
//       };
//       this.lexer.inline(token.text, token.tokens); // Queue this data to be processed for inline tokens blockTokens
//       return token;
//     }
//   },
//   renderer(token) {
//     return `<div class='tip' style="color:red">${this.parser.parseInline(token.tokens)}</div>`; // parseInline to turn child tokens into HTML
//   },
// };

// const description = {
//   name: 'description',
//   level: 'inline',                                 // Is this a block-level or inline-level tokenizer?
//   start(src) { return src.match(/:/)?.index; },    // Hint to Marked.js to stop and check for a match
//   tokenizer(src, tokens) {
//     const rule = /^:([^:\n]+):([^:\n]*)(?:\n|$)/;  // Regex for the complete token, anchor to string start
//     const match = rule.exec(src);
//     if (match) {
//       return {                                         // Token to generate
//         type: 'description',                           // Should match "name" above
//         raw: match[0],                                 // Text to consume from the source
//         dt: this.lexer.inlineTokens(match[1].trim()),  // Additional custom properties, including
//         dd: this.lexer.inlineTokens(match[2].trim())   //   any further-nested inline tokens
//       };
//     }
//   },
//   renderer(token) {
//       return `\n<div class='tip' style="color:red">${this.parser.parseInline(token.dd)}</div>`
//     // return `\n<dt>${this.parser.parseInline(token.dt)}</dt><dd>${this.parser.parseInline(token.dd)}</dd>`;
//   },
//   childTokens: ['dt', 'dd'],                 // Any child tokens to be visited by walkTokens
// };

const descriptionList = {
  name: 'descriptionList',
  level: 'block',                                     // Is this a block-level or inline-level tokenizer?
  start(src) { 
   // console.log( src.match(/:tip:\n/)?.index,"start",src);
    return src.match(/:tip:/)?.index;
     }, // Hint to Marked.js to stop and check for a match /:tip:[^:\n]/
  tokenizer(src, tokens) {
    const rule = /^(?::tip:)[\S\s]*?(?<=::\n)/;    // Regex for the complete token, anchor to string start /^(?::[^:\n]+:[^:\n]*(?:\n|$))+/
    const match = rule.exec(src);
  console.log("match",match,src);
    if (match) {
      const token = {                                 // Token to generate
        type: 'descriptionList',                      // Should match "name" above
        raw: match[0],                                // Text to consume from the source
        text: match[0].trim(),                        // Additional custom properties
        tokens: []                                    // Array where child inline tokens will be generated
      };
      this.lexer.inline(token.text, token.tokens); 
      //this.lexer.inline(token.text, token.tokens);    // Queue this data to be processed for inline tokens  blockTokens
      return token;
    }
  },
  renderer(token) {
    return `<div class='tip' style="color:red">${this.parser.parseInline(token.tokens)}</div>`; // parseInline to turn child tokens into HTML 
  }
};
 
marked.use({ extensions: [descriptionList] }); 

const rawHtml = marked.parse(`
水水水水
:tip:
Topic 1:  Description 1   
:Topic 1:  Description 1  
::

## 混合

:tip:
Topic 1:  Description 1   
:Topic 1:  Description 1  
::

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

<style scoped>
.read-the-docs {
  color: #888;
}
</style>

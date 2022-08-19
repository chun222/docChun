<!--
 * @Date: 2022-08-10 15:20:28
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-19 13:58:13
 * @FilePath: \pages\src\components\layout\footer.vue
-->
<template>
<div>
      <a-space> <span>当前文档由ChunDoc创建</span> <a href="https://github.com/chun222/docChun" target="_blank"><github theme="outline" size="24" fill="#fff"/></a> </a-space>
</div>
    
</template>

<style scoped>

</style>

<script>
import {
 Github
} from "@icon-park/vue-next";
export default {
    components:{
        Github
    },
    setup() {
        
    },
}
</script>
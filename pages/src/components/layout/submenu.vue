<!--
 * @Date: 2022-08-10 16:13:42
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-30 14:19:16
 * @FilePath: \pages\src\components\layout\submenu.vue
-->
<template>
    <a-sub-menu :key="menuInfo.fullpath"> 
      <template #title>{{ menuInfo.name }}</template>
      <template v-for="item in menuInfo.children" :key="item.fullpath">
        <template v-if="!item.children">
          <a-menu-item :key="item.fullpath">  
          <router-link :to="{ name: dataParams.routeName, params: {project: dataParams.project, lang: dataParams.lang,version:dataParams.version, page:item.relativepath}}">{{
            item.name
          }}</router-link>
          </a-menu-item>
        </template>
        <template v-else>
          <sub-menu :menu-info="item" :key="item.fullpath" />
        </template>
      </template>
    </a-sub-menu>
</template>
<style scoped>

</style>

<script lang="ts">  
export default {
  name: 'SubMenu',
  props: {
    dataParams:{
      type: Object,
    },
    menuInfo: {
      type: Object,
      default: () => ({}),
    },
  },
   components: { 
  }, 
};
</script>
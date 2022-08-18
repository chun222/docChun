<!--
 * @Date: 2022-08-10 16:13:42
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-18 10:30:45
 * @FilePath: \pages\src\components\layout\submenu.vue
-->
<template>
    <a-sub-menu :key="menuInfo.fullpath"> 
      <template #title>{{ menuInfo.name }}</template>
      <template v-for="item in menuInfo.children" :key="item.fullpath">
        <template v-if="!item.children">
          <a-menu-item :key="item.fullpath"> 
           <router-link
            :to="`/${encodeURIComponent(item.fullpath)}`" 
            >{{ item.name }}</router-link
          >
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
    menuInfo: {
      type: Object,
      default: () => ({}),
    },
  },
   components: { 
  }, 
};
</script>
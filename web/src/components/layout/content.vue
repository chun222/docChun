<!--
 * @Date: 2021-11-30 15:42:07
 * @LastEditors: 春贰
 * @Desc: 
 * @LastEditTime: 2022-08-11 09:50:35
 * @FilePath: \web\src\components\layout\content.vue
-->
<template>
  <!-- 框架 content 区域 -->
  <div id="content">
    <router-view v-slot="{ Component, route }" v-if="true">
      <transition :name="route.meta.transition || 'fade'" mode="out-in">
        <!-- 客户端永远更新 -->
        <keep-alive v-if="true">
          <component :is="Component" :key="route.name"></component>
        </keep-alive>
        <component v-else :is="Component" :key="route.name"></component>
      </transition>
    </router-view>
    <div class="content-loading" v-else>
      <a-spin size="large" />
    </div>
  </div>
</template>
<script>
import { computed } from "vue";
export default {
  setup() {
    const keepAlive = true;
    return {
      keepAlive,
    };
  },
};
</script>
<style  scoped>

</style>

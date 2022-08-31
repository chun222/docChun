<!--
 * @Date: 2021-11-30 15:42:07
 * @LastEditors: 春贰
 * @Desc: 
 * @LastEditTime: 2022-08-31 14:05:16
 * @FilePath: \pages\src\components\layout\content.vue
-->
<template>
  <!-- 框架 content 区域 -->
  <router-view v-slot="{ Component, route }" v-show="!Loading">
    <transition :name="'fade'" mode="out-in">
      <!-- <transition :name="route.meta.transition || 'fade'" mode="out-in"> -->
        <div>
          
      <component :is="Component" :key="route.name"></component>
        </div>

      <!-- <keep-alive v-if="true">
          <component :is="Component" :key="route.name"></component>
        </keep-alive>
        <component v-else :is="Component" :key="route.name"></component> -->
    </transition>
  </router-view>

  <div class="content-loading" v-if="Loading">
    <a-spin size="large" />
  </div>
</template>
<script lang="ts">
import { computed } from "vue";
import { useStore } from "@/store/index";
export default {
  setup() {
    const store = useStore();
    const Loading = computed(() => store.routeLoading);
    const keepAlive = true;
    return {
      keepAlive,
      Loading,
    };
  },
};
</script>
<style  scoped>
</style>

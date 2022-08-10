<!--
 * @Date: 2022-08-10 16:13:42
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-10 16:20:42
 * @FilePath: \web\src\components\layout\submenu.vue
-->
<template>
    <a-sub-menu :key="menuInfo.key">
      <template #icon><MailOutlined /></template>
      <template #title>{{ menuInfo.title }}</template>
      <template v-for="item in menuInfo.children" :key="item.key">
        <template v-if="!item.children">
          <a-menu-item :key="item.key">
            <template #icon>
              <PieChartOutlined />
            </template>
            {{ item.title }}
          </a-menu-item>
        </template>
        <template v-else>
          <sub-menu :menu-info="item" :key="item.key" />
        </template>
      </template>
    </a-sub-menu>
</template>
<style scoped>

</style>

<script> 
import { 
  PieChartOutlined,
  MailOutlined,
} from '@ant-design/icons-vue';
export default {
  name: 'SubMenu',
  props: {
    menuInfo: {
      type: Object,
      default: () => ({}),
    },
  },
   components: {
    PieChartOutlined,
    MailOutlined,
  }, 
};
</script>
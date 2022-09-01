<!--
 * @Date: 2022-09-01 08:44:46
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-09-01 17:27:56
 * @FilePath: \pages\src\components\layout\drawer\menus.vue
-->
<template>
  <a-drawer
    :width="500"
    title="目录配置"
    :placement="placement"
    :visible="visible"
    @close="onClose"
    :getContainer="parentNode"
  >
    <template #extra>
      <a-button style="margin-right: 8px" @click="onClose">取消</a-button>
      <a-button type="primary" @click="onSubmit">提交</a-button>
    </template>
    <div class="configmenutree">
      <a-row>
        <a-col span="24">
          <a-form ref="formRef" :model="formState" autocomplete="off">
            <a-input :value="formState.parent" type="hidden" />

            <a-form-item
              label="上级目录"
              name="parentname" 
            >
              <a-input
                :value="formState.parentname ? formState.parentname : '根目录'"
                placeholder="请选择左侧菜单"
                readonly
              />
            </a-form-item>

            <a-form-item
              label="目录或文件名"
              name="dir"
              :rules="[{ required: true, message: '必须填写' }]"
            >
              <a-input
                v-model:value="formState.dir"
                placeholder="请选择左侧菜单"
              />
            </a-form-item>

            <a-form-item
              label="显示名称"
              name="name"
              :rules="[{ required: true, message: '必须填写' }]"
            >
              <a-input
                v-model:value="formState.name"
                placeholder="请选择左侧菜单"
              />
            </a-form-item>

            <a-form-item
              label="类型"
              name="type"
              :rules="[{ required: true, message: '必须填写' }]"
            >
              <a-select
                v-model:value="formState.type"
                :disabled="formState.typereadonly"
              >
                <a-select-option :value="'dir'"> 目录</a-select-option>
                <a-select-option :value="'md'"> 文件</a-select-option>
              </a-select>
            </a-form-item>

            <a-form-item
              label="排序"
              name="position"
              :rules="[{ required: true, message: '必须填写' }]"
            >
              <a-input-number
                v-model:value="formState.position"
                placeholder="请选择左侧菜单"
              />
            </a-form-item>
          </a-form>
        </a-col>

        <a-col span="24">
          <a-divider orientation="left">目录列表</a-divider>
          <a-tree
            @select="selectTree"
            :autoExpandParent="true"
            :tree-data="treeData"
          >
            <template #title="{ fullpath, name, type }">
              <a-dropdown :trigger="['contextmenu']">
                <span>{{ name }}</span>
                <template #overlay>
                  <a-menu
                    @click="
                      ({ key: menuKey }) =>
                        onContextMenuClick(fullpath, name, menuKey)
                    "
                  >
                    <a-menu-item v-if="type == 'dir'" :key="1">
                      <add-three theme="outline" /> 目录下添加</a-menu-item
                    >
                    <a-menu-item :key="2"
                      ><delete theme="outline" /> 删除</a-menu-item
                    >
                  </a-menu>
                </template>
              </a-dropdown>
            </template>
          </a-tree>
        </a-col>
      </a-row>
    </div>
  </a-drawer>
</template>

<script lang="ts">
import {
  computed,
  defineComponent,
  reactive,
  ref,
  watch,
  ComputedRef,
  toRaw,
  getCurrentInstance,
} from "vue";
import type { DrawerProps, FormInstance } from "ant-design-vue";
import { AddThree, Delete } from "@icon-park/vue-next";
import { useStore } from "@/store/index";
import { AliasDirType } from "@/model";
import { saveconfigs } from "@/api/module/base";

export default defineComponent({
  components: {
    AddThree,
    Delete,
  },
  props: {
    visible: {
      type: Boolean,
    },
    parentNode: {
      type: HTMLElement,
    },
  },
  emits: ["close"],
  setup(props, { emit }) {
    const globleConfig = getCurrentInstance()?.proxy;
    const placement = ref<DrawerProps["placement"]>("right");

    const formRef = ref<FormInstance>();
    const formState = reactive({
      name: "",
      path: "",
      dir: "",
      position: 0,
      type: "dir",
      parent: "",
      parentname: "",
      typereadonly: true,
    });

    const onSubmit = () => {
      formRef.value.validate().then(() => {
            console.log(formState);

      });
    };
    const onClose = () => {
      emit("close");
    };

    const onContextMenuClick = (
      fullpath: string,
      name: string,
      menuKey: string | number
    ) => {
      console.log(`fullpath: ${fullpath},name: ${name}, menuKey: ${menuKey}`);
      if (menuKey == 1) {
        formState.parent = fullpath;
        formState.parentname = name;
        formState.typereadonly = false
      } else if (menuKey == 2) {
        //删除
      }
    };

    const selectTree = (v, v2) => {
        
        formState.typereadonly = true
      formState.path = v2.node.fullpath;
      formState.name = v2.node.name;
      formState.dir = v2.node.realname;
      formState.type = v2.node.type;
      formState.position = v2.node.position;
      formState.parent = v2.node.parent ? v2.node.parent.node.fullpath : "";
      formState.parentname = v2.node.parent ? v2.node.parent.node.name : "";
    };

    const store = useStore();

    const treeData = computed(() => store.menus);

    const setbaseDir = () => {
      formState.parent = "";
      formState.parentname = "";
    };

    return {
      placement,
      onClose,
      onSubmit,
      onContextMenuClick,
      treeData,
      setbaseDir,
      formState,
      formRef,
      selectTree,
    };
  },
});
</script>



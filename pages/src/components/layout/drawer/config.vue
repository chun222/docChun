<!--
 * @Date: 2022-09-01 08:44:46
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-09-01 14:35:24
 * @FilePath: \pages\src\components\layout\drawer\config.vue
-->
<template>
  <a-drawer
    :width="500"
    title="系统配置"
    :placement="placement"
    :visible="visible"
    @close="onClose"
    :getContainer="parentNode"
  >
    <template #extra>
      <a-button style="margin-right: 8px" @click="onClose">取消</a-button>
      <a-button type="primary" @click="onSubmit">提交</a-button>
    </template>

    <a-form ref="formRef" name="dynamic_form_nest_item" :model="allconfigs">
      <!-- begin -->
      <a-divider orientation="left">项目配置</a-divider>
      <a-space
        v-for="(project, index) in allconfigs.projects"
        :key="project.id"
        style="display: flex; margin-bottom: 8px"
        align="baseline"
      >
        <a-form-item
          label="文件夹"
          :name="['projects', index, 'dir']"
          :rules="{
            required: true,
            message: '必须填写',
          }"
        >
          <a-input
            v-model:value="project.dir"
            placeholder="文件夹"
            :readonly="project.readonly !== false"
          />
        </a-form-item>

        <a-form-item
          label="名称"
          :name="['projects', index, 'name']"
          :rules="{
            required: true,
            message: '必须填写',
          }"
        >
          <a-input v-model:value="project.name" placeholder="项目名称" />
        </a-form-item>
        <delete theme="outline" @click="removeproject(project)" />
      </a-space>
      <a-form-item>
        <a-button type="dashed" block @click="addproject">
          <add-three theme="outline" />
          新增项目
        </a-button>
      </a-form-item>
      <!-- 项目end -->

      <!-- begin -->
      <a-divider orientation="left">版本配置</a-divider>
      <a-space
        v-for="(project, index) in allconfigs.versions"
        :key="project.id"
        style="display: flex; margin-bottom: 8px"
        align="baseline"
      >
        <a-form-item
          label="文件夹"
          :name="['versions', index, 'dir']"
          :rules="{
            required: true,
            message: '必须填写',
          }"
        >
          <a-input
            v-model:value="project.dir"
            placeholder="文件夹"
            :readonly="project.readonly !== false"
          />
        </a-form-item>

        <a-form-item
          label="名称"
          :name="['versions', index, 'name']"
          :rules="{
            required: true,
            message: '必须填写',
          }"
        >
          <a-input v-model:value="project.name" placeholder="版本名称" />
        </a-form-item>
        <delete theme="outline" @click="removeversion(project)" />
      </a-space>
      <a-form-item>
        <a-button type="dashed" block @click="addversion">
          <add-three theme="outline" />
          新增版本
        </a-button>
      </a-form-item>
      <!-- 项目end -->

      <!-- begin -->
      <a-divider orientation="left">多语言配置</a-divider> 
      <a-space
        v-for="(project, index) in allconfigs.langs"
        :key="project.id"
        style="display: flex; margin-bottom: 8px"
        align="baseline"
      >
        <a-form-item
          label="文件夹"
          :name="['langs', index, 'dir']"
          :rules="{
            required: true,
            message: '必须填写',
          }"
        >
          <a-input
            v-model:value="project.dir"
            placeholder="文件夹"
            :readonly="project.readonly !== false"
          />
        </a-form-item>

        <a-form-item
          label="名称"
          :name="['langs', index, 'name']"
          :rules="{
            required: true,
            message: '必须填写',
          }"
        >
          <a-input v-model:value="project.name" placeholder="项目名称" />
        </a-form-item>
        <delete theme="outline" @click="removelang(project)" />
      </a-space>
      <a-form-item>
        <a-button type="dashed" block @click="addlang">
          <add-three theme="outline" />
          新增语言
        </a-button>
      </a-form-item>
      <!-- 项目end -->
    </a-form>
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
import {AliasDirType} from '@/model';
import {saveconfigs} from '@/api/module/base'; 

export default defineComponent({
  components: {
    AddThree,
    Delete,
  },
  props: {  
    visible: {
      type: Boolean,
    }, 
    parentNode:{
      type:HTMLElement
    }
  },
  emits: ["close"],
  setup(props, { emit }) {
    const globleConfig = getCurrentInstance()?.proxy;
    const placement = ref<DrawerProps["placement"]>("right");

    const formRef = ref<FormInstance>();

    const onSubmit = () => {
      formRef.value.validate().then(() => { 
        //重载
        saveconfigs(allconfigs).then((re)=>{
          if (re.code==0) { 
             store.InitConfig().then(()=>{
                  globleConfig.$notice.success({message:"保存成功"})
             })
          }else{
            globleConfig.$notice.error({message:re.msg})
          }
        })
      });
    };
    const onClose = () => {
      emit("close");
    };

    const store = useStore();

    const allconfigs = reactive<{
      projects: ComputedRef<AliasDirType[]>;
      versions: ComputedRef<AliasDirType[]>;
      langs: ComputedRef<AliasDirType[]>;
    }>({
      projects:computed(()=> store.projects) ,
      versions: computed(()=> store.versions),
      langs:computed(()=> store.langs),
    });

    // 项目
    const removeproject = (item: AliasDirType) => {
       if (allconfigs.projects.length<=1) {
          globleConfig.$message.error("必须保留一项")
          return
      }
      let index = allconfigs.projects.indexOf(item);
      if (index !== -1) {
        allconfigs.projects.splice(index, 1);
      }
    };
    const addproject = () => {
      allconfigs.projects.push({
        dir: "",
        name: "",
        readonly: false,
        id: Date.now(),
      });
    };

    // 项目
    const removeversion = (item: AliasDirType) => {
       if (allconfigs.versions.length<=1) {
          globleConfig.$message.error("必须保留一项")
          return
      }
      let index = allconfigs.versions.indexOf(item);
      if (index !== -1) {
        allconfigs.versions.splice(index, 1);
      }
    };
    const addversion = () => {
      allconfigs.versions.push({
        dir: "",
        name: "",
        readonly: false,
        id: Date.now(),
      });
    };

    // 项目
    const removelang = (item: AliasDirType) => {
      if (allconfigs.langs.length<=1) {
          globleConfig.$message.error("必须保留一项")
          return
      }
      let index = allconfigs.langs.indexOf(item);
      if (index !== -1) {
        allconfigs.langs.splice(index, 1);
      }
    };
    const addlang = () => {
      allconfigs.langs.push({
        dir: "",
        name: "",
        readonly: false,
        id: Date.now(),
      });
    };

   
    return {
      placement,
      onClose,
      onSubmit,
      formRef,
      allconfigs,
      removeproject,
      addproject,
      removeversion,
      addversion,
      removelang,
      addlang,  
    };
  },
});
</script>



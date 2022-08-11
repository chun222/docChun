<template>
 <a-modal
    :visible="visible"
    :title="t('info')"
    :cancelText="t('cancel')"
    :okText="t('submit')"
    @ok="submit"
    @cancel="cancel(false)"
  >
    <a-form
       ref="formRef"
      :model="record"
      :rules="formRules"
      :label-col="labelCol"
      :wrapper-col="wrapperCol"
    > 
      {{range .Fields }}
       <a-form-item ref="{{.Json}}" :label="t('{{.Comment}}')" name="{{.Json}}">
        <a-input disabled v-model:value="record.{{.Json}}" />
      </a-form-item> 
    {{end}}  
    </a-form>
  </a-modal>
</template>
<script>
import { useI18n  } from "vue-i18n";
import { message } from 'ant-design-vue';
import { add, tree } from "@/api/module/permission";
import { defineComponent, reactive, ref, toRaw, watch } from "vue";
export default defineComponent({
  props: {
    visible: {
      type: Boolean,
    },
    record: {
      type: Object
    }
  },
  emit: ["close"],
  setup(props, context) {
    const { t } = useI18n();
    const state = reactive({ 
      depts: [], 
    })
    
    const formRef = ref();
    
    const formState = reactive({ 
    });

    watch(props,(props) => { 
    })

    const formRules = { 
    };

    const submit = (e) => {
      formRef.value.resetFields();
      context.emit("close", false);
    };

    const cancel = (e) => {
      formRef.value.resetFields();
      context.emit("close", false);
    };
 
    return {
      t,
      state,
      submit,
      cancel,
      formRef,
      formState,
      formRules, 
      labelCol: { span: 6 },
      wrapperCol: { span: 18 }, 
      fieldNames: {children:'children', title:'title', key:'id', value: 'id' }
    };
  },
});
</script>
<template>
  <a-modal
    :visible="visible"
    :title="t('edit')"
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
     {{range .Fields }}{{if eq .InputType "input"}} 
       <a-form-item ref="{{.Json}}" :label="t('{{.Comment}}')" name="{{.Json}}">
        <a-input v-model:value="record.{{.Json}}" />
      </a-form-item>
     {{else if eq .InputType "textarea"}} 
     <a-form-item ref="{{.Json}}" :label="t('{{.Comment}}')" name="{{.Json}}">
        <a-textarea v-model:value="record.{{.Json}}" />
      </a-form-item>
      {{else if eq .InputType "switch"}} 
     <a-form-item ref="{{.Json}}" :label="t('{{.Comment}}')" name="{{.Json}}">
        <a-switch :checkedValue="1" :unCheckedValue="0" :checkedChildren="t('enable')" :unCheckedChildren="t('disalbe')" v-model:checked="record.{{.Json}}" />
      </a-form-item>
      {{else if eq .InputType "select"}} 
     <a-form-item ref="{{.Json}}" :label="t('{{.Comment}}')" name="{{.Json}}">
        <a-select v-model:value="record.{{.Json}}">
          <a-select-option value="v1"> field1 </a-select-option>
          <a-select-option value="v2"> field2 </a-select-option> 
        </a-select>
      </a-form-item>
    {{ end }}{{end}}
    </a-form>
  </a-modal>
</template>
<script>
import { useI18n  } from "vue-i18n";
import { message } from 'ant-design-vue';
import { edit } from "@/api/module/{{.TableName}}";
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
      showComponent: true,
      showPath: true,
      showCode: false
    });
    
    const formRef = ref();
    
    let formState = reactive({ 
    });

    const formRules = {
      {{range .Fields }}{{if .Notnull}}{{if ne .InputType "none"}} {{.Json}}: [{ required: true, message: t('please input first'), trigger: "blur" }],
      {{end}}{{end}}{{ end }}  
    };

    const submit = (e) => {
      formRef.value
        .validate()
        .then(() => {
          edit(toRaw(props.record)).then((response)=>{
              if(response.code== 0 ){
                message.success({ content: t('Saved successfully!'), duration: 1 }).then(()=>{
                  cancel(true);
                });
              }else{
                message.error({ content: response.msg, duration: 1 }).then(()=>{
                  cancel(false);
                });
              }
          });
        })
        .catch(error => {
          console.log('error', error);
        });
    };

     
    const cancel = (e=false) => {
      formRef.value.resetFields();
      context.emit("close", e);
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
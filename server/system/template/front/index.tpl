<template> 
    <page-layout>
      <a-row :gutter="[10, 10]">
        <!-- 顶部区域 -->
        <a-col :span="24">
          <a-card>
            <!-- 查询参数 -->
            <go2-query
              :searchParam="searchParam"
              @on-search ="search"
            ></go2-query>
          </a-card>
        </a-col>
        <!-- 中心区域 -->
        <a-col :span="24">
          <a-card>
            <!-- 列表 -->
            <go2-table
              ref="tableRef"
              rowKey="id"
              :indentSize="10"
              :fetch="fetch"
              :columns="columns"
              :toolbar="toolbar"
              :operate="operate"
              :param="state.param"
              :pagination="pagination"
              :row-selection="{ selectedRowKeys: state.selectedRowKeys, onChange: onSelectChange }"
            >
              <!-- 继承至 a-table 的默认插槽 -->
            </go2-table>
          </a-card>
        </a-col>
      </a-row>
      <add v-if="state.visibleAdd" :visible="state.visibleAdd"  @close="closeAdd"></add>
      <edit v-if="state.visibleEdit" :visible="state.visibleEdit" @close="closeEdit" :record="state.recordEdit"></edit>
      <info v-if="state.visibleInfo" :visible="state.visibleInfo" @close="closeInfo" :record="state.recordInfo"></info>
    </page-layout>
</template>

<script>
import { useI18n  } from "vue-i18n";
import add from "./modal/add.vue";
import edit from "./modal/edit.vue";
import info from "./modal/info.vue";
import { message , Modal} from 'ant-design-vue';
import { ExclamationCircleOutlined } from '@ant-design/icons-vue';
import { list, remove } from "@/api/module/{{.TableName}}";
import { reactive, createVNode, ref,toRaw } from 'vue';

const removeKey = "remove"; 

export default {
  components: {
    add,
    edit,
    info,
  },
  setup() {
    const { t } = useI18n();
    const tableRef = ref();

    const switchFormat = { yes: 1, no: 0, event: function(value,record){ 
      return value;
    }}

    /// 删除配置
    const removeMethod = (record) => {
       removeIds([record.id])
    }

    //批量删除
    const batchremoveMethod = () => { 
      const romovekeys =  toRaw(state.selectedRowKeys) ;
      if (romovekeys.length<=0) {
        message.error({content:t('please chose'), key: removeKey, duration: 1})
        return false;
      }
       removeIds(romovekeys)
    }

    const removeIds = (Ids) => {
      Modal.confirm({
        title: t('confirm')+t('delete') +'?',
        icon: createVNode(ExclamationCircleOutlined),
        okText: t('ok'),
        cancelText: t('cancel'),
        onOk() {
          message.loading({ content: t('loading'), key: removeKey }); 
          remove({"ids":Ids}).then((response) => {
            if(response.code==0){
              message.success({content: t('Delete succeeded!'), key: removeKey, duration: 1}).then(()=> {
                tableRef.value.reload()
              })
            }else{
              message.error({content:response.msg, key: removeKey, duration: 1})
            }
          })
        }
      });
    }

     

    /// 列配置
    const columns = [
     {{range .Fields }} 
     {{if .List}}{{if eq .Type "time.Time"}}{title:"{{.Comment}}", dataIndex: "{{.Json}}", key: "{{.Json}}", format: { type: "date", value: "YYYY-MM-DD HH:mm:ss" } },{{else}}{ title:"{{.Comment}}",dataIndex: "{{.Json}}", key: "{{.Json}}" },{{end}}{{end}}
    {{ end }} 
    ];

    /// 数据来源  
    const fetch = async (param) => {  
      var response = await list(param);
      if (response.code!=0) { 
         message.error({content:response.msg, duration: 1}) 
      } 
       return {data :response.data }
     
    };

    /// 工具栏
    const toolbar = [
      { label: 'add', code: '/{{.TableName}}/add', event: function () { state.visibleAdd = true }},
      { label: 'batchRemove', code: '/{{.TableName}}/remove', event:  batchremoveMethod }
    ];

    /// 行操作 
    const operate = [
      { label: 'info', code:'/{{.TableName}}/list', event: function (record) { state.visibleInfo = true, state.recordInfo = record }},
      { label: 'edit', code:'/{{.TableName}}/edit', event: function (record) { state.visibleEdit = true, state.recordEdit = record }},
      { label: 'remove', code:'/{{.TableName}}/remove', event: function (record) { removeMethod(record) }},
    ];

     /// 分页参数
    const pagination = reactive({
      pageSize: 20,
      page: 1,
    });
    
    /// 外置参数 - 当参数改变时, 重新触发 fetch 函数
    const state = reactive({
      selectedRowKeys: [],
      param: {},
      visibleAdd: false,
      visibleEdit: false,
      visibleInfo: false,
      recordEdit: {},
      recordInfo: {},
      
    }) 
    /// 查询参数
    //tips:自动生成的下拉需要手动设置
    const searchParam = [
    {{range $k, $v :=.SearchFields }}
    {{ if eq $v.InputType "switch"}}
    { key: "{{$v.Json}}", type: "select", label: '{{$v.Comment}}', value: "",
          hidden: {{ge $k 3}} ,
          options: [
            { text: 'v1', value: "v1"},
            { text: 'v2', value: "v2"}, 
          ]
        },{{ else if eq $v.InputType "select"}}
    { key: "{{$v.Json}}", type: "select", label: '{{$v.Comment}}', value: "",
          hidden: {{ge $k 3}} ,
          options: [
            { text: 'v1', value: "v1"},
            { text: 'v2', value: "v2"}, 
          ]
        },{{else}}{ key: "{{$v.Json}}", type: "input", label: '{{$v.Comment}}',hidden:{{ge $k 3}} },{{ end }}{{ end }}  
    ]
 

    /// 查询操作
    const search = function(value) {
      state.param = value 
    }

    const onSelectChange = selectedRowKeys => { 
      state.selectedRowKeys = selectedRowKeys;
    };

    const closeAdd = function(e){
      state.visibleAdd = false;  
      if (e) {
        tableRef.value.reload()
      }
    }

    const closeEdit = function(e){
      state.visibleEdit = false; 
      if (e) {
        tableRef.value.reload()
      }
      
    }

    const closeInfo = function(e){
       state.visibleInfo = false;
       if (e) {
        tableRef.value.reload()
      }
    }

    return {
      t,
      state, 
      fetch, 
      search,
      toolbar, 
      columns, 
      operate, 
      pagination,
      searchParam,
      onSelectChange,

      closeAdd,
      closeEdit,
      closeInfo,

      tableRef
    };
  },
};
</script>
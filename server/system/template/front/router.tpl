export default {  
  '/{{.TableName}}/list': () => import('@/appview/{{.TableName}}/index.vue'),  
}
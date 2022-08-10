/*
 * @Date: 2022-06-02 16:18:46
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-06-02 16:48:50
 * @FilePath: \web\src\tools\pdf.js
 */
import { genpdf } from "@/api/module/tools"; 
import { message } from "ant-design-vue";
 

export const exportPdf= (name,html,format="A4",pageSize="Portrait") =>  new Promise((resolve, reject) => { //导出pdf
  genpdf({ name: name, html: html,format:format,pageSize:pageSize }).then(
   (res) => { 
     if (typeof res.msg != "undefined") {
       message.error(res.msg);
       return reject(err.data); 
     } 
     /*pdf处理*/ 
     let blob = new window.Blob([res], {
       type: "application/pdf;charset=UTF-8",
     });
     const href = URL.createObjectURL(blob);
     window.open(href);
     return resolve();
   }
 );
})
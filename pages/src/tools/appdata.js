/*
 * @Date: 2022-04-12 10:36:43
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-04-12 10:41:16
 * @FilePath: \web\src\tools\appdata.js
 */

import { sequence } from "@/api/module/tools";

export const getSequence = async (pre,length) => {  //获取序列号
   return  sequence({pre:pre,length:length})   
}
/*
 * @Date: 2022-08-15 08:42:35
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-30 12:05:48
 * @FilePath: \pages\src\tools\common.ts
 */ 


import router from "@/route/index";
import { doclist } from "@/api/module/base";
import { useStore } from "@/store/index";
export const getMenu =  async (isReload: boolean) =>  {
  const store = useStore();
  const dataParams: any = {};
  dataParams.project = store.project.dir;
  dataParams.lang = store.lang.dir;
  dataParams.version = store.version.dir;
 
 await doclist(dataParams).then((re) => {
    if (re.code == 0) {
      //默认指向第一篇文章
      //isEmptyObjValue(router.currentRoute.value.params)
      //  const isIndex = isEmptyObjValue(router.currentRoute.value.params);
      //是否有页面
      const isNotHaspage = router.currentRoute.value.params.page == "";
 
      store.menus = re.data;
      if (re.data.length > 0) {
        if (isNotHaspage || isReload) {
          dataParams.page = re.data[0].fullpath;
        } else {
          dataParams.page = router.currentRoute.value.params.page;
        }
        store.page = dataParams.page ;  
        if (router.currentRoute.value.name=="/") {
          router.push({
            name: "/",
            params: dataParams,
          });
        }
      }
    }
  });

  return dataParams ; 
};


//复制文本
export function copyText(text:string) {
   // 添加一个input元素放置需要的文本内容
   const copyContent= document.createElement('input');
   copyContent.value = text;
   document.body.appendChild(copyContent);
   // 选中并复制文本到剪切板
   copyContent.select();
   document.execCommand('copy'); 
   // 移除input元素
   document.body.removeChild(copyContent);
  
}


//防抖  
export function debounce(func:Function, wait = 500, immediate = false ) {
  let timeout;
  return function () {
    let context = this, args = arguments;
    let later = function () {
      timeout = null;
      if (!immediate) func.apply(context, args);
    };
    let callNow = immediate && !timeout;
    clearTimeout(timeout);
    timeout = setTimeout(later, wait);
    if (callNow) func.apply(context, args);
  };
}

//节流
type TypethrottleOp = {
  leading ?:boolean ; 
  trailing ?:boolean ; 
}
export function throttle(func:Function, wait = 500, options?:TypethrottleOp) { 
  let context, args, result;
  let timeout = null;
  let previous = 0;
  if (!options) options = {};
  let later = function () {
    previous = options.leading === false ? 0 : Date.now();
    timeout = null;
    result = func.apply(context, args);
    if (!timeout) context = args = null;
  };
  return function () {
    let now = Date.now();
    if (!previous && options.leading === false) previous = now;
    let remaining = wait - (now - previous);
    context = this;
    args = arguments;
    if (remaining <= 0 || remaining > wait) {
      if (timeout) {
        clearTimeout(timeout);
        timeout = null;
      }
      previous = now;
      result = func.apply(context, args);
      if (!timeout) context = args = null;
    } else if (!timeout && options.trailing !== false) {
      timeout = setTimeout(later, remaining);
    }
    return result;
  };
}

//生成uuid
export function uuid() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
    var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
    return v.toString(16);
  });
}
 

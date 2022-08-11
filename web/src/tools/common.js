
import { onUnmounted } from "vue";
import store from "@/store"  
/**
 * 刷新路由
 */
export const refresh = async () => { 
  store.commit("app/UPDATE_ROUTER_ACTIVE");
  setTimeout(() => {
    store.commit("app/UPDATE_ROUTER_ACTIVE");
  }, 500);
};

/**
 * 开发环境
 */
export const isNotProduction = () => {
  return import.meta.env.NODE_ENV !== 'production'
}

/**
 * 定时时间
 * 
 * @param timeout   超时事件
 * @param callback  回调事件
 */
export const isTimeout = (timeout, callback) => {
  setTimeout(() => {
    callback();
  }, timeout.value * 1000);
  const Interval = setInterval(() => {
    timeout.value--;
  }, 1000);
  onUnmounted(() => {
    clearInterval(Interval);
  });
}

/**
 * 根据 当前路径 查询 所有父级 (包括当前)
 * 
 * @param arr 菜单列表
 * @param id 指定路由
 */
 export const findParentAll = (arr, id) => {
  var temp = []
  var forFn = function (list, path) {
    for (var i = 0; i < list.length; i++) {
      var item = list[i]
      if (item.path === path) {
        temp.push(findPathById(arr, item.id))
        forFn(list, item.parent);
        break
      } else {
        if (item.children) {
          forFn(item.children, path)
        }
      }
    }
  }
  forFn(arr, id)
  return temp
}

/**
 * 获取图片地址 
 *  
 * @param name  name
 */
 export const getImgUrl = (name) => {
  return new URL(`/src/assets/image/${name}`, import.meta.url).href
}
 
/**
 * 根据 当前路径 查询 所有父级 
 * 
 * @param arr 菜单列表
 * @param id  当前路由
 */
export const findParent = (arr, id) => {
  var temp = []
  var forFn = function (list, path) {
    for (var i = 0; i < list.length; i++) {
      var item = list[i]
      if (item.path === path) {
        temp.push(findPathById(arr, item.parent))
        forFn(list, item.parent);
        break
      } else {
        if (item.children) {
          forFn(item.children, path)
        }
      }
    }
  }
  forFn(arr, id)
  return temp
}

/**
 * 根据 当前路径 查询 菜单编号
 * 
 * @param arr 菜单列表
 * @param id  当前路由
 */
export const findPathById = (arr, key) => {
  var forFn = function (list, id) {
    for (var i = 0; i < list.length; i++) {
      var item = list[i]
      if (item.id === id) {
        return item.path;
      } else {
        if (item.children) {
          forFn(item.children, name)
        }
      }
    }
  }
  return forFn(arr, key);
}

export function resetChartStyle() {
   if(store.getters.clientMode=="admin"){
    let c_width = document.documentElement.clientWidth;
    let c_height = document.documentElement.clientHeight;
    let zoom_W = c_width / 1920;
    let zoom_H = c_height / 1080;
    let zoom_act = zoom_W > zoom_H ? zoom_H : zoom_W;
    return {
        zoom: 1 / zoom_act,
        transform: "scale(" + zoom_act + ")",
        transformOrigin: "0%0%",
        width: 1 / zoom_act * 100 + "%"
    };
   }else{
     return {}
   }
 
}


export function showLoading() {
  return {
      text: '加载中',
      color: '#1BC85E',
      textColor: '#ccc',
      maskColor: 'rgba(255, 255, 255, 0.0)',
      zlevel: 0
  }
}



//转换色彩
export function hexToRgba(hex, opacity) {
  return 'rgba(' + parseInt('0x' + hex.slice(1, 3)) + ',' + parseInt('0x' + hex.slice(3, 5)) + ','
      + parseInt('0x' + hex.slice(5, 7)) + ',' + opacity + ')';
}

//判断是否是图片
export function isImg(typename) {
  typename = typename.toLowerCase()
   return  typename=='png' || typename=='jpg' ||  typename=='jpeg' ||  typename=='gif' ||  typename=='ico' ||  typename=='tif' ||  typename=='webp' 
}


import { message } from "ant-design-vue";
//复制文本
export function copyText(text) {
   // 添加一个input元素放置需要的文本内容
   const copyContent= document.createElement('input');
   copyContent.value = text;
   document.body.appendChild(copyContent);
   // 选中并复制文本到剪切板
   copyContent.select();
   document.execCommand('copy'); 
   // 移除input元素
   document.body.removeChild(copyContent);
   message
   .success({
     content: "copy succeeded" , 
     duration: 1,
   }) 
}


//防抖  
export function debounce(func, wait = 500, immediate = false ) {
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
export function throttle(func, wait = 500, options = {}) { 
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
 

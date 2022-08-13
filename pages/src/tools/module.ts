/*
 * @Date: 2022-08-12 23:24:32
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-08-13 22:08:58
 * @FilePath: \pages\src\tools\module.ts
 */
/**
 * 路径拼接
 * 
 * @param 根路径
 * @param 子路径 
 */
 export const resolve = (root:string, path:string) => {
    if (path.startsWith('/')) {
        return root + path;
    } else {
        return root + "/" + path;
    }
}

/**
 * 模块化导入 
 * 
 * @param context 
 */
export const module = (context:any) => {
    return Object.keys(context).reduce((modules, key) => {
        return {
            ...modules,
            ...context[key].default
        }
    }, {})
}
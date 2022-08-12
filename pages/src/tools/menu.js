/*
 * @Date: 2021-11-30 15:42:07
 * @LastEditors: 春贰
 * @Desc: 
 * @LastEditTime: 2022-07-08 10:56:10
 * @FilePath: \web\src\tools\menu.js
 */
 
 
export const toTree = (list) => {  
    const data = JSON.parse(JSON.stringify(list)) // 浅拷贝不改变源数据
    const result = []
    if (!Array.isArray(data)) {
      return result
    }
    data.forEach(item => {
      delete item.children
    })
    const map = {}
    data.forEach(item => {
      map[item.id] = item
    })
    data.forEach(item => {
      const parent = map[item.parent]
      if (parent) {
        (parent.children || (parent.children = [])).push(item)
      } else {
        result.push(item)
      }
    })

    return result
  
} 

/**
 * 路由是否注册 
 * 
 * @param routes 路由集合
 * @param currentPath 当前路由
 */
export const hasRoute = (routes, currentPath) => {
    let boolean = false
    for (let i = 0; i < routes.length; i++) {
        const { name, children = [] } = routes[i]
        if (name === currentPath) {
            boolean = true
        } else if (children.length > 0) {
            boolean = hasRoute(children, currentPath)
        }
        if (boolean) {
            break
        }
    }
    return boolean;
}
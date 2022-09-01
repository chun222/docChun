/*
 * @Date: 2022-09-01 11:28:27
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 
 * @LastEditTime: 2022-09-01 11:32:43
 * @FilePath: \pages\src\model\index.ts
 */
export type TypeDoclist = {
    project: string;
    lang: string;
    version: string;
}

export type TypePath = {
    project: string;
    lang: string;
    version: string;
    page: string
    content?: string
}


export type TypeSearch = {
    project: string;
    lang: string;
    version: string;
    keyword: string;
}

export interface AliasDirType {
    name: string;
    dir: string;
    id?: number;
    readonly?: boolean;
}

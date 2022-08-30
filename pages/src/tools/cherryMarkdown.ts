//配置
import Cherry from "cherry-markdown";
import CherryEngine from 'cherry-markdown/dist/cherry-markdown.engine.core';
import { pinyin } from "pinyin-pro";
import echarts from "echarts";
import { svgArr } from "@/view/index/svg";
// 自定义菜单
const AddPrefixTemplate = Cherry.createMenuHook("AddPrefixTemplate", {
    iconName: "icon-add-prefix",
    onClick(selection: string) {
        return "Prefix-" + selection;
    },
    shortcutKeys: [], //快捷键集合, 用于注册键盘函数，当匹配的快捷键组合命中时，也会调用click函数
    //子菜单
    //name: String
    // iconName： String
    // noIcon: Boolean
    // onClick: Function
    subMenuConfig: [],
});


const CustomHookA = Cherry.createSyntaxHook('codeBlock', Cherry.constants.HOOKS_TYPE_LIST.PAR, {
    makeHtml(str) {
      console.warn('custom hook', 'hello');
      return str;
    },
    rule(str) {
      const regex = {
        begin: '',
        content: '',
        end: '',
      };
      regex.reg = new RegExp(regex.begin + regex.content + regex.end, 'g');
      return regex;
    },
  });

const cherryEngineInstance = new CherryEngine(null) as any;  
/*
 *  tips Hook
    语法 Hook 类型，仅可选 SEN（行内语法）、PAR（段落语法）
 */
const BlockTipsHook = Cherry.createSyntaxHook(
    "blockTipsHook",
    Cherry.constants.HOOKS_TYPE_LIST.PAR,
    { 
        // 开启缓存，用于保护内容
       needCache: true,
       // 预处理文本，避免受影响,生命周期，返回替换后的字符串
        beforeMakeHtml(str: string) {
            return str.replace(this.RULE.reg, (match, code) => {
 
                //获取匹配内部内容
                const content =  cherryEngineInstance.makeHtml(match.replace(/^:(info|success|warning|error):[\s\n]|::$/g, ""))

                const lineCount = (match.match(/\n/g) || []).length;
                const sign = this.$engine.md5(match);
                const html = `<div data-sign="${sign}" data-lines="${lineCount + 1
                    }" ><div class="${code}" >
                    <div class="tip-header">${svgArr[code] || ""} <span>${code}</span></div>
                    <p> ${content}</p>
                    </div></div>`;
                   
                return this.pushCache(html, sign, lineCount);
 
            });
        },
        //处理成html之后的
        makeHtml(str: string) { 
             return str
        },
        rule(str: string) {
            return {
                // reg: /(?<=:(info|success|warning|error):[\s\n])[\S\s]*?(?=::)/g,
                reg: /:(info|success|warning|error):[\s\n][\S\s]*?::/g,
            };
        },
    }
);


/*
 * 生成一个屏蔽敏感词汇的hook
 * 名字叫blockSensitiveWords
 * 匹配规则会挂载到实例的RULE属性上
 */
const BlockSensitiveWordsHook = Cherry.createSyntaxHook(
    "blockSensitiveWords",
    Cherry.constants.HOOKS_TYPE_LIST.PAR,
    {
        // 开启缓存，用于保护内容
        needCache: true,
        // 预处理文本，避免受影响
        beforeMakeHtml(str) {
            return str;
            // return str.replace(this.RULE.reg, (match, code) => {
            //     const lineCount = (match.match(/\n/g) || []).length;
            //     const sign = this.$engine.md5(match);
            //     const html = `<div data-sign="${sign}" data-lines="${lineCount + 1
            //         }" >***</div>`;
            //     return this.pushCache(html, sign, lineCount);
            // });
        },
        makeHtml(str, sentenceMakeFunc) {
            return str;
        },
        rule(str) {
            return {
                reg: /sensitive words/g,
            };
        },
    }
);
export const BasicConfig = {
    id: "markdown",
    fileUpload(file, callback) {
        console.log("上传了：", file);
        callback("");
    },
    externals: {
        echarts: echarts,
        // katex:katex,
        MathJax: (window as any).MathJax,
    },
    isPreviewOnly: false,
    engine: {
        global: {
            urlProcessor(url, srcType) {
                console.log(`url-processor`, url, srcType);
                return url;
            },
        },
        syntax: {
            codeBlock: {
                theme: 'twilight',
              },
            codeBlock2: {
                theme: "dark", // 默认为深色主题
                wrap: true, // 超出长度是否换行，false则显示滚动条
                lineNumber: true, // 默认显示行号
                customRenderer: {
                    // 自定义语法渲染器
                },
                /**
                 * indentedCodeBlock是缩进代码块是否启用的开关
                 *
                 *    在6.X之前的版本中默认不支持该语法。
                 *    因为cherry的开发团队认为该语法太丑了（容易误触）
                 *    开发团队希望用```代码块语法来彻底取代该语法
                 *    但在后续的沟通中，开发团队发现在某些场景下该语法有更好的显示效果
                 *    因此开发团队在6.X版本中才引入了该语法
                 *    已经引用6.x以下版本的业务如果想做到用户无感知升级，可以去掉该语法：
                 *        indentedCodeBlock：false
                 */
                indentedCodeBlock: false,
            },
            table: {
                enableChart: false,
                // chartEngine: Engine Class
            },
            fontEmphasis: {
                allowWhitespace: false, // 是否允许首尾空格
            },
            strikethrough: {
                needWhitespace: false, // 是否必须有前后空格
            },
            mathBlock: {
                engine: "MathJax", // katex或MathJax
                src: "https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-svg.js", // 如果使用MathJax plugins，则需要使用该url通过script标签引入
            },
            inlineMath: {
                engine: "MathJax", // katex或MathJax
            },
            emoji: {
                useUnicode: false,
                customResourceURL:
                    "https://github.githubassets.com/images/icons/emoji/unicode/${code}.png?v8",
                upperCase: true,
            },
            toc: {
                /** 默认只渲染一个目录 */
                allowMultiToc: false,
            },
            header: {
                /**
                 * 标题的样式：
                 *  - default       默认样式，标题前面有锚点
                 *  - autonumber    标题前面有自增序号锚点
                 *  - none          标题没有锚点
                 */
                anchorStyle: "autonumber",
            },
        },
        customSyntax: {
            CustomHook: {
                syntaxClass: CustomHookA,
                force: false,
                after: 'br',
              },
            // 注入编辑器的自定义语法中
            BlockSensitiveWordsHook: {
                syntaxClass: BlockSensitiveWordsHook,
                // 有同名hook则强制覆盖
                force: true,
                // 在处理图片的hook之前调用
                // before: 'image',
            },
            BlockTipsHook: {
                syntaxClass: BlockTipsHook,
                // 有同名hook则强制覆盖
                force: true,
            }
        },
    },
    toolbars: {
        toolbar: [
            "addPrefix",
            "bold",
            "italic",
            "strikethrough",
            "|",
            "color",
            "header",
            "ruby",
            "|",
            "list",
            {
                insert: [
                    "image",
                    "audio",
                    "video",
                    "link",
                    "hr",
                    "br",
                    "code",
                    "formula",
                    "toc",
                    "table",
                    "pdf",
                    "word",
                    "ruby",
                ],
            },
            "graph",
            "togglePreview",
            "settings",
            "switchModel",
            "codeTheme",
            "export",
        ],
        bubble: [
            "bold",
            "italic",
            "underline",
            "strikethrough",
            "sub",
            "sup",
            "quote",
            "ruby",
            "|",
            "size",
            "color",
        ], // array or false
        sidebar: ["mobilePreview", "copy"],
        customMenu: {
            // 注入编辑器的菜单中
            // 对象 key 可以作为菜单项的名字（需要保证唯一），在上方的配置中使用
            addPrefix: AddPrefixTemplate,
        },
    },
    editor: {
        codemirror: {
            // depend on codemirror theme name: https://codemirror.net/demo/theme.html
            // 自行导入主题文件: `import 'codemirror/theme/<theme-name>.css';`
           //  theme: "default",
        },
        // 编辑器的高度，默认100%，如果挂载点存在内联设置的height则以内联样式为主
        height: "100%",
        // defaultModel 编辑器初始化后的默认模式，一共有三种模式：1、双栏编辑预览模式；2、纯编辑模式；3、预览模式
        // edit&preview: 双栏编辑预览模式
        // editOnly: 纯编辑模式（没有预览，可通过toolbar切换成双栏或预览模式）
        // previewOnly: 预览模式（没有编辑框，toolbar只显示“返回编辑”按钮，可通过toolbar切换成编辑模式）
        defaultModel: "edit&preview",
        // 粘贴时是否自动将html转成markdown
        convertWhenPaste: true,
    },
    previewer: {
        dom: false, 
        // 是否启用预览区域编辑能力（目前支持编辑图片尺寸、编辑表格内容）
        enablePreviewerBubble: true,
    },
    keydown: [],
    //extensions: [],
    callback: {
        changeString2Pinyin: pinyin,
    },
};
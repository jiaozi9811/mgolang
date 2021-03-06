# markdown

[TOC]

## TOC

使用TOC添加目录[TOC]

## 换行

1. 连续两个以上空格+回车方法
2. 使用html语言换行标签：\<br>

## 缩进

不断行的空白格`&nbsp;`或`&#160;`
半角的空格`&ensp;`或`&#8194;`
全角的空格`&emsp;`或`&#8195;`

## 页眉 页脚

~页眉

@页脚

## 粗体 斜体 删除线

*斜体*
_斜体_
**粗体**
**_粗体 斜体_**
***粗体 斜体***
~~删除线~~

## 背景色

==背景色==

## 引用

>引用内容

>如果引用内容需要换行  
>可以在行尾添加两个空格

>或者在引用内容中加一个空行

## 分隔线

***
<!--下划线-->
___
---

## 代码块

```javascript
代码块
代码块
代码块  
代码块  
```

`代码块`

## 锚点 行内链接

[markdown](#markdown)

## 表情符号

:smiley

:smirk

<https://www.webpagefx.com/tools/emoji-cheat-sheet/>

## 列表

### 无序列表

* 可以使用 `*` 作为标记
* 可以使用 `*` 作为标记


+ 也可以使用 `+`
+ 也可以使用 `+`


- 或者 `-`
- 或者 `-`

### 有序列表

1. 有序列表以数字和 `.` 开始；
3. 数字的序列并不会影响生成的列表序列；
4. 但仍然推荐按照自然顺序（1.2.3...）编写。

### 嵌套列表

1. 第一层
    + 1-1
    + 1-2
2. 无序列表和有序列表可以随意相互嵌套
    1. 2-1
    2. 2-2

## 选择列表

- [ ] Eat
- [x] Code
    - [x] HTML
    - [ ] CSS
    - [x] JavaScript
- [ ] Sleep

## 表格

name | age
-|-
LearnShare | 12
Mike |  32

### 对齐

:--- 代表左对齐
:--: 代表居中对齐
---: 代表右对齐

## 超链接

[Google](http://www.google.com/ "Google")

[icon.png](./images/icon.png)

<http://www.google.com/>

## 图片

![image](https://static.segmentfault.com/v-5c78d357/global/img/creativecommons-cc.svg)

![image][9]

[9]:https://static.segmentfault.com/v-5c78d357/global/img/creativecommons-cc.svg

指定大小

<img src="https://avatars2.githubusercontent.com/u/3265208?v=3&s=100" alt="GitHub" title="GitHub,Social Coding" width="50" height="50" />

## LaTex数学公式

1. 行内公式：使用两个”$”符号引用公式
2. 行间公式：使用两对“$$”符号引用公式

## 流程图

## 折叠details

<details open>
  <summary>点击时的区域标题：点击查看详细内容</summary>
  <p> - 测试 测试测试</p>
    <pre><code>  title，value，callBack可以缺省  </code>  </pre>
</details>

- summary：折叠语法展示的摘要
- details：折叠语法标签
- pre：以原有格式显示元素内的文字是已经格式化的文本
- blockcode：表示程序的代码块
- code：指定代码范例

## 注脚

脚注是在需要标记脚注文字的后面增加一个方括号，方括号中的内容必须以^开头，再接着是数字、字符串标记

接着，在文件的任意地方，你可以把这个脚注的内容定义出来

注脚^[注脚内容], 注脚^[注脚内容]
注脚^[注脚内容]

注脚^[注脚内容]

注脚^[注脚内容]

注脚^[*]

[^*]: something
# Contributing to SHUPT-GO

## What should I know before I get started?

### 技术栈

我们使用Go语言和Beego框架。

数据库使用Postgres（通过beego/orm）。

将会使用到GraphQL。

## Styleguides

### Git使用

#### Git Commit Messages

- Commit message的第一行长度不超过30个字

- Commit message要表达出这次commit做了什么

- Commit message可以用下面的表情来开头：
  - 🎨 `:art:` 表示整理代码、重构
  - 📝 `:memo:` 表示编写文档
  - 🐛 `:bug:` 表示修复一个bug
  - 🔥 `:fire:` 表示移除文件（夹）
  - 🔒 `:lock:` 表示处理与安全有关的问题
  - ⬆️ `:arrow_up:` 表示升级依赖
  - ⬇️ `:arrow_down:` 表示降级依赖
  - ✅ `:white_check_mark:` 表示添加测试
  - :hammer_and_wrench:`:hammer_and_wrench:` 表示添加功能

#### Git commit 和 push的时机

如果你能做到这些，那你可以commit and push一次：

- 确保你的代码在你能想到的范围内能够工作
- 确保你能为本次commit写一个有意义的Commit Message

在commit前请删去所有注释掉的代码，如果某天你真的需要这些注释掉的代码，你可以用版本控制工具找回。

同理，请不要在commit的代码中包含用于调试的输入输出语句。

#### 当你的代码和线上代码冲突时

使用变基（Rebase）而非合并（Merge）来处理冲突。

### Go language Styleguide

⚠️：你可以拒绝遵守这些规定，但请你在commit了这么做的代码之后，开一个issue来找一个愿意这么做的人来帮你修改你的代码。

请在提交代码前用`gofmt`（默认配置）对代码进行格式化。

### 变量、函数和自定义类型的命名

- 使用小驼峰（eg. camelCase）来为局部变量、类型和函数命名。

- 使用大驼峰（eg. CamelCase）来为导出的自定义类型和函数命名。

- 命名时的自然语言使用英文，**绝对禁止**混用英文和中文拼音。

- 变量名要反应变量代表什么，类似的，类型名要反应类型是什么，函数名要反应函数做了什么。

- 变量、函数和类型的名称尽量保持在9~16个字符之间

  - 除非在“变量是循环索引”或“变量代表数学上的变量”的情况下，否则**绝对禁止**使用单个字符的变量名。
  - 除非是数学上约定好的变量名，否则**绝对禁止**使用仅由数字区别的几个变量名（eg. x1,x2）。

- 为相同的东西取相同的名字，为不同的东西取不同的名字，为相反的东西取相反的名字

  - 一些标准反义词

    | 原义   | 反义     |
    | ------ | -------- |
    | add    | remove   |
    | begin  | end      |
    | create | destory  |
    | first  | last     |
    | insert | delete   |
    | get    | set      |
    | lock   | unlock   |
    | min    | max      |
    | next   | previous |
    | open   | close    |
    | show   | hide     |
    | start  | stop     |
    | up     | down     |

- **如果变量名太长**，请按顺序考虑以下这些方法

  1. 去掉虚词（and、or、the）

  2. 使用标准缩写，一些缩写如下：
  
       | 全称     | 缩写 |
       | -------- | ---- |
       | index    | idx  |
       | object   | obj  |
       | document | doc  |
       | text     | txt  |
       | position | pos  |

  3. 去掉所有非前置元音（但保证单词仍然能被拼读出来）（eg. computer -> cmptr，screen -> scrn）

       如果你使用了这一条，请将你使用的缩写添加到上面的标准缩写表中

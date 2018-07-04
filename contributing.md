# Contributing to SHUPT-GO

## What should I know before I get started?

### 技术栈

我们使用Go语言和Beego框架。

数据库使用Postgres（通过beego/orm）。

**考虑**使用GraphQL。

## Styleguides

### Git Commit Messages

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

  ## Go language Styleguide

  请在提交代码前用`gofmt`（默认配置）对代码进行格式化
![image-20240402094632607](https://sse-market-source-1320172928.cos.ap-guangzhou.myqcloud.com/blog/image-20240402094632607.png)

# 软工集市，软工人定义的世界

### 项目介绍

SSE_MARKET是一个跨校区的学院内部交流平台，主要以论坛的形式为软工师生提供自由交流和信息汇集的半匿名空间，可在上面看帖、搜索、发帖、评论、回复、点赞等，提供宽松、平等的交流氛围，致力于消除学院内部的信息差以及促进数字文化传承。

SSE_MARKET又称软工集市，主要由21级本科生组成的 SSE_MARKET小组负责设计、开发、部署和维护，它脱胎于软件中级实训课，并在学院的支持下逐步发展，现注册加入师生约 400人，学生包括大一到大四本科生以及各级研究生，教师包括行政老师、技术老师到专业老师等。

[软工集市成功交接，欢迎新任成员 – SSE_MARKET博客 (ssemarket.cn)](https://ssemarket.cn/2024/04/02/软工集市成功交接，欢迎新任成员/)

在2024.3.6，软工集市正式完成交接。现在的软工集市由22级和23级本科生组成的新小组负责开发、部署、优化、维护。

[官方博客](https://ssemarket.cn)

# 管理员端介绍

## 运营背景与功能概述

为了满足运营组在实际工作中的需求，我们特别开发了管理员端。该端旨在为运营团队提供高效便捷的管理工具，以确保平台内容的优质与规范。主要功能包括帖子的增删改查、优质帖子的设定、邀请码的管理等。

### 帖子管理

管理员可以通过直观的界面完成帖子的创建、删除、修改和查询操作。无论是日常内容维护还是处理违规内容，都能迅速响应。此外，还设有专门的筛选功能，可按时间、类别等多种维度查看帖子，方便进行精细化管理。

### 优质帖子设定

为了鼓励高质量内容的创作与传播，管理员有权将优秀帖子标记为 “优质帖子”。被标记的帖子将在前端展示时获得特殊标识，吸引更多用户关注与互动。

### 邀请码管理

邀请码是控制新用户注册节奏的重要手段。管理员可在后台批量生成、分配、启用或禁用邀请码。


### 如何运行软工集市管理系统后端

SSEMARKET后端为go项目，需要配置go开发的基本环境。此外，还需要配置mysql和redis的数据库环境。

1. 首先需要在开发环境安装go语言，理论上安装最新版本就行。
   [Download and install - The Go Programming Language](https://golang.google.cn/doc/install)

2. 选择自己喜欢的IDE配置开发环境，这里推荐vscode,goland

3. 安装go项目依赖

4. 安装配置mysql8.0

5. 在`config/application.yml`改数据库配置

6. ```shell
   go run main.go
   ```

   输入该命令即可运行项目。

注意事项：1、2 、4、5主要是在配置基本环境，如果本来配置好了可以直接跳过。
这里并没有给出具体的安装配置步骤，而是直接贴出了参考教程，这个一方面是由于Windows、MacOS和linux等操作系统的配置过程会有所不同，因此给出教程也不一定适用（比如mysql，redis的教程是适用windows的）。


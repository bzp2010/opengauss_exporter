---
title: 软件架构
---

## 架构图 {#architecture-graph}
![image](/img/concepts/architecture-graph-cn.png)

## 架构 {#architecture}

### 模块信息 {#architecture-module}

#### 刮削管理器 {#architecture-module-scrape}
刮削任务管理器负责集中管理多个刮削任务（即每一个数据库节点），统一控制启停。

#### 刮削器 {#architecture-module-scraper}
每一个刮削器都是一个具有专用场景的程序模块，在数据库中执行SQL语句，并将返回的数据进行处理后存入缓存中。

#### 内存缓存 {#architecture-module-cache}
在Go语言中实现内存缓存模块，将 Exporter 和刮削器解耦，提供更稳定的 Metrics 输出性能

#### 服务器 {#architecture-module-server}
提供HTTP服务器，实现 Prometheus SDK 中的 Exporter 接口，通过 `Collect` 函数调用从内存缓存中调取数据生成 Metrics API 响应
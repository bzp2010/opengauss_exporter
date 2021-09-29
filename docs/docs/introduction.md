---
title: 简介
slug: "/"
---

## openGauss 数据库 {#opengauss}

openGauss是一款开源的关系型数据库,采用客户端/服务器，单进程多线程架构，支持单机和一主多备部署方式，备机可读，支持双机高可用和读扩展。

### 主要特点 {#opengauss-characteristic}

#### 高性能 {#opengauss-characteristic-performance}

提供了面向多核架构的并发控制技术结合鲲鹏硬件优化，在两路鲲鹏下TPCC Benchmark达成性能150万tpmc。
针对当前硬件多核numa的架构趋势， 在内核关键结构上采用了Numa-Aware的数据结构。
提供Sql-bypass智能快速引擎技术。

#### 高可用 {#opengauss-characteristic-availability}

支持主备同步，异步以及级联备机多种部署模式。
数据页CRC校验，损坏数据页通过备机自动修复。
备机并行恢复，10秒内可升主提供服务。

#### 高安全 {#opengauss-characteristic-security}

支持全密态计算，访问控制、加密认证、数据库审计、动态数据脱敏等安全特性，提供全方位端到端的数据安全保护。

#### 易运维 {#opengauss-characteristic-maintenance}

基于AI的智能参数调优和索引推荐，提供AI自动参数推荐。
慢SQL诊断，多维性能自监控视图，实施掌控系统的性能表现。
提供在线自学习的SQL时间预测。

#### 全开放 {#opengauss-characteristic-open}

采用木兰宽松许可证协议，允许对代码自由修改，使用，引用。
数据库内核能力全开放。
提供丰富的伙伴认证，培训体系和高校课程。
openGauss相比其他开源数据库主要有多存储模式，NUMA化内核结构和高可用等产品特点。

## Prometheus {#prometheus}

Prometheus 是由 Soundcloud 以开源软件的形式进行发布的监控和告警软件，许多公司和组织都采用了 Prometheus 作为其监控告警工具。Prometheus 于 2016 年 5 月加入 CNCF 基金会，成为继 Kubernetes 之后的第二个 CNCF 托管项目。
Prometheus 主要包括 Server 和时序数据库两部分，Server负责对外提供查询、写入API，时序数据库负责以时间顺序存储采集来的 Metric 指标数据。

### Metric {#prometheus-metric}

Prometheus采集到的监控数据均以 Metric（指标）形式保存在时序数据库中（TSDB）

#### Metric 指标格式
```text
<metric name>{<label key>=<label value>, ...}
```
- metric name: 监控指标名称
- label key = label value: 标签键值  
通过 label 系统可以标记数据（eg. server=127.0.0.1:5432）

#### Metric 指标类型
- Counter 计数值，累加数据
- Gauge 计量值，可增可减
- Histogram 直方图
- Summary 摘要

### Server

Prometheus Server 直接从监控目标中或者间接通过推送网关来拉取监控指标，并在本地存储所有抓取到的样本数据。

### Exporter

在 Prometheus Server 中 Pull 模式使用的组件。它从目标服务中提取 Metric 指标数据，并将其组装为 Prometheus 可识别的格式，通过对外开放的 API 为 Prometheus 刮削器采集存储。

## openGauss Exporter

openGauss Exporter 是为 openGauss 数据库实现的 Prometheus Exporter，它通过 SQL 从数据库节点中读取运行环境指标等信息，并组装成 Prometheus Metric 格式，向外提供采集接口。

它可以同时连接多个 openGauss 节点，以一定的时间间隔采集运行数据。这些数据将被临时存储在内存缓存中，以供采集 API 调用。

:::note 
当前仅测试 openGauss 2.0.1 版本兼容性
:::
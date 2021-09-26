---
title: 5分钟体验
---

## 概述
本文是 openGauss 的快速入门指南的快速体验部分。快速入门分为两个步骤：
1. 通过 Docker 安装：
    - openGauss 数据库
    - openGauss Exporter
    - Prometheus Server
2. 等待片刻，即可通过 Prometheus UI 查看指标数据

## 前提条件
- 已安装 Docker 的运行环境

## 第一步：安装各个组件

1. 安装 openGauss 数据库
```docker
docker run --name opengauss --privileged=true -d -e GS_PASSWORD=gaussdb!123 enmotech/opengauss:latest
```
2. 创建 openGauss Exporter 配置文件
```yaml
server:
  http:
    host: 0.0.0.0
    port: 9188

data_sources:
  - dsn: "postgresql://gaussdb:gaussdb!123@127.0.0.1:5432/postgres"
    duration: 5s
    max_retry: 3
    master: true
    enable_postgresql_exporter: true
    enable_settings: true
    enable_os_run_info: true
    enable_total_memory_detail: true
    enable_sql_count: true
    enable_instance_time: true
```
复制以上配置文件存储成名为`config.yaml`的文件

3. 启动 openGauss Exporter
```docker
docker run -d -v $(pwd)/config.yaml:/etc/opengauss_exporter.yaml -p 9188:9188 bzp2010/opengauss_exporter -c /etc/opengauss_exporter.yaml
```

4. 创建 Prometheus 配置文件
```yaml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: "prometheus"
    static_configs:
      - targets: ["localhost:9090"]

  - job_name: "opengauss"
    scrape_interval: 5s
    static_configs:
      - targets: ["127.0.0.1:9188"]
```
复制以上配置文件存储成名为`prometheus.yaml`的文件

5. 启动 Prometheus Server
```docker
docker run -d -p 9090:9090 -v $(pwd)/prometheus.yaml:/etc/prometheus/prometheus.yml prom/prometheus
```

## 第二步：访问 Prometheus UI 并查看指标
1. 等待Exporter初次刮削及Prometheus Server初次采集，约需15s
2. 使用浏览器访问 `127.0.0.1:9090`
3. 使用 PromQL 查询指标 `{server:"127.0.0.1:5432"}`
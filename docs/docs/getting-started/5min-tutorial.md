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

1. 将本仓库克隆至本地
```shell
git clone https://github.com/bzp2010/opengauss_exporter
cd opengauss_exporter
```
2. 进入体验目录
```shell
cd example\tutorial
```

3. 使用 Docker Compose 启动体验环境
```shell
docker-compose up -d
```

4. 等待1-2分钟

全新的 openGauss 数据库节点初次运行会进行数据库初始化，需要等待数分钟

5. 进行简单测试
```shell
curl http://127.0.0.1:9188/
```

## 第二步：访问 Prometheus UI 并查看指标
1. 使用浏览器访问 `127.0.0.1:9090`
2. 使用 PromQL 查询指标 `{server:"127.0.0.1:5432"}`
---
title: HTTP 接口
---

## GET `/metrics`

输出 Prometheus 格式的指标数据

## GET `/refresh`

清除 Metrics 内存缓存数据

- 响应
  - 成功：refresh success
  - 失败：refresh failed: {失败原因}
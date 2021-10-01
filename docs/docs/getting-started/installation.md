---
title: 安装
---

## Docker {#docker}
```shell
$ docker pull bzp2010/opengauss_exporter:latest
$ docker run -d bzp2010/opengauss_exporter:latest
```

:::note关于镜像版本
当前可能并无正式版本发布，因此请使用dev镜像进行测试，即 `bzp2010/opengauss_exporter:dev`
:::

## Source {#source}
Required: Golang **1.14 +**

```shell
$ git clone https://github.com/bzp2010/opengauss_exporter.git
$ cd opengauss_exporter
$ go build -o opengauss_exporter main.go
$ ./opengauss_exporter help
```
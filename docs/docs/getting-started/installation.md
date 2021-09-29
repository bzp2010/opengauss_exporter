---
title: 安装
---

## Docker {#docker}
```shell
$ docker pull bzp2010/opengauss_exporter:v1.0.0
$ docker run -d bzp2010/opengauss_exporter:v1.0.0
```

## Source {#source}
required: Golang 1.14 +

```shell
$ git clone https://github.com/bzp2010/opengauss_exporter.git
$ cd opengauss_exporter
$ go build -o opengauss_exporter main.go
$ ./opengauss_exporter
```
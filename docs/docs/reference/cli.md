--- 
title: 命令行
---

## 总览 {#summary}
```text
A OpenGauss metric exporter for Prometheus

Usage:
  opengauss_exporter [flags]
  opengauss_exporter [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  serve       run openGauss Exporter
  version     Print the version information of Exporter

Flags:
  -c, --config string   config file path (default "config.yaml")
  -h, --help            help for opengauss_exporter
```

## 参数 {#flags}
- `-c` 允许用户绕过默认配置文件地址，指定自定义路径作为配置文件地址

## 命令 {#command}
### help
查看程序命令行帮助信息

### completion
生成特定命令的 shell 脚本

### serve
启动 Exporter 程序服务器，并开始连接数据库刮削数据

### version
查看当前 Exporter 的版本、编译时间、编译分支Hash等信息
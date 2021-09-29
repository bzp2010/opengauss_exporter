---
title: 配置文件
---

```yaml
server:               # server configure
  http:               # HTTP server
    host: 0.0.0.0     # host eg. 127.0.0.1
    port: 9188        # port eg. 9188
  https:              # HTTPS server
    host: 0.0.0.0     # host eg. 127.0.0.1
    port: 9189        # port eg. 9189
    cert: "/etc/ssl/cert.pem"   # cert file path
    key: "/etc/ssl/private.key" # private key file path

data_sources:      # data source configure
  - dsn: "postgresql://gaussdb:gaussdb@127.0.0.1:5432/postgres"
    duration: 5s   # scrape duration eg. 3s 5m 36h 1d12h
    max_retry: 3   # max connect retries
    master: true   # is master node (Not supported temporary)
    
    # enable or disable scraper
    enable_settings: true             # pg_settings
    enable_os_run_info: true          # gs_os_run_info
    enable_total_memory_detail: true  # gs_total_memory_detail
    enable_sql_count: true            # gs_sql_count
    enable_instance_time: true        # gs_instance_time
    enable_postgresql_exporter: true  # postgresql exporter

  - dsn: "postgresql://gaussdb:gaussdb@127.0.0.1:5433/postgres"
    duration: 8s
    max_retry: 3
    master: true
    enable_settings: true
    enable_os_run_info: true
    enable_total_memory_detail: true
    enable_sql_count: true
    enable_instance_time: true
    enable_postgresql_exporter: true

```
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

tasks:    # scrape task configure
  - dsn: "postgresql://gaussdb:gaussdb@127.0.0.1:5432/postgres"
    name: Test Server 1    # scrape task name (only for label)
    duration: 5s           # scrape duration eg. 3s 5m 36h 1d12h
    master: true           # is master node (master-slave architecture)
    scrapers:              # enabled scraper (ALL scraper)
      - postgresql_exporter     # cover PostgreSQL Exporter metrics
      - pg_settings             # pg_settings view
      - gs_os_run_info          # gs_os_run_info view
      - gs_instance_time        # gs_instance_time view
      - gs_total_memory_detail  # gs_total_memory_detail view
      - gs_sql_count            # gs_sql_count view

  - dsn: "postgresql://gaussdb:gaussdb@127.0.0.1:5433/postgres"
    name: Test Server 2
    duration: 5s
    master: true
    scrapers:
      - postgresql_exporter
      - pg_settings
      - gs_os_run_info
      - gs_instance_time
      - gs_total_memory_detail
      - gs_sql_count
```
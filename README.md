# openGauss Server Exporter
Prometheus exporter for openGauss server metrics.

CI Tested openGauss versions: `2.0.1`

## Building and running
> 1. install Golang ***Above 1.14***   
> 2. move config.yaml.example to config.yaml and edit it   
> 3. run these commands   

    git clone https://github.com/bzp2010/opengauss_exporter.git
    cd opengauss_exporter
    go build -o opengauss_exporter main.go
    ./opengauss_exporter serve -c config.yaml

## Edit Config

```
server:
  http:
    host: 0.0.0.0
    port: 9188       # metrics API endpoints

data_sources:
  - dsn: "postgresql://gaussdb:gaussdb@127.0.0.1:5432/postgres"
    duration: 5s     # scrape metrics interval
    max_retry: 3     # retry times on connect failed
    master: true     # NOTICE temporarily useless
    enable_settings: true             # pg_settings view
    enable_os_run_info: true          # gs_os_run_info view
    enable_total_memory_detail: true  # gs_total_memory_detail view
    enable_sql_count: true            # gs_sql_count view
    enable_instance_time: true        # gs_instance_time view
    enable_postgresql_exporter: true  # this scraper come from postgresql_exporter 
```
## About metrics
### gs_settings_*
These metrics are openGauss config items.

### gs_os_run_info_*
These metrics are current database instance's OS info.

### gs_instance_time_*
These metrics are current database instance used times on any part.

### gs_sql_count_*
These metrics are executed times of every type of SQL.

### gs_total_memory_detail
These metrics are memory usage information.

#### And more metrics will be added.

## TODO List
- [ ] Add docker support for deploy
- [ ] Add Grafana Dashboard template
- [ ] Add more metrics for OS status
- [ ] Add more E2E test case
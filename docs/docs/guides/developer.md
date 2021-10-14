---
title: 开发者指引
---

## 刮削器扩展指南 {#developer-scraper}

### 创建刮削器并实现 `scrape.Scraper` 接口 {#developer-scraper-struct}

:::note路径
internal/scraper/xxx.go
:::

```go
type BaseInfoScraper struct {
}

func NewBaseInfoScraper() *BaseInfoScraper {
    return &BaseInfoScraper{}
}

func (g BaseInfoScraper) Scrape(t *scrape.TaskOLD) ([]prometheus.Metric, []error, error) {
    return []prometheus.Metric{}, nil, nil
}
```

### 在框架中注册这个刮削器 {#developer-scraper-register}

:::note路径
internal/scraper/scraper.go
:::

```go
func Init()  {
    // ....
    scrape.RegisterScraper("base_info", NewBaseInfoScraper())
    // ....
}
```

### 在配置文件中启用注册的刮削器  {#developer-scraper-config}

```yaml
tasks:
  - dsn: "postgresql://gaussdb:gaussdb@127.0.0.1:5432/postgres"
    name: Test Server 1
    duration: 5s
    master: true
    scrapers:
      - base_info  # 使用在上一步中注册的刮削器ID，即可启动已注册的刮削器
```
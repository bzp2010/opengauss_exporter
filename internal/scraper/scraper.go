package scraper

import (
	"github.com/prometheus/client_golang/prometheus"

	"opengauss_exporter/internal/core/exporter"
	"opengauss_exporter/internal/core/scrape"
)

func Init()  {
	scrape.RegisterScraper("postgresql_exporter", NewBuiltinSQLScraper())
	scrape.RegisterScraper("pg_settings", NewPgSettingsScraper())
	scrape.RegisterScraper("gs_os_run_info", NewGsOSRunInfoScraper())
	scrape.RegisterScraper("gs_instance_time", NewGsInstanceTimeScraper())
	scrape.RegisterScraper("gs_total_memory_detail", NewGsTotalMemoryDetailScraper())
	scrape.RegisterScraper("gs_sql_count", NewGsSQLCountScraper())
}

func newDesc(subsystem string, name string, help string, labels prometheus.Labels, hiddenNamespace... bool) *prometheus.Desc {
	namespace := exporter.NamespaceView
	if len(hiddenNamespace) > 0 && hiddenNamespace[0] {
		namespace = ""
	}

	return prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, name),
		help, nil, labels,
	)
}
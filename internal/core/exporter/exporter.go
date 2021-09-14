package exporter

import (
	"github.com/prometheus/client_golang/prometheus"

	"opengauss_exporter/internal/core/cache"
	"opengauss_exporter/internal/utils"
)

type Exporter struct{}

var (
	Name = "opengauss_exporter"

	Namespace = "gs"

	SubsystemScrapeTask = "scrape_task"

	LabelServer = "server"
)

func NewExporter() *Exporter {
	return &Exporter{}
}

func (e Exporter) Describe(_ chan<- *prometheus.Desc) {
}

func (e Exporter) Collect(ch chan<- prometheus.Metric) {
	data, err := cache.Metrics.Data()
	if err != nil {
		utils.GetLogger().Error("Get cache data failed", err)
		return
	}

	for _, v := range data {
		metrics := v.([]prometheus.Metric)
		for _, metric := range metrics {
			ch <- metric
		}
	}
}

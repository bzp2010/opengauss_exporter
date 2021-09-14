package scraper

import (
	"github.com/prometheus/client_golang/prometheus"

	"opengauss_exporter/internal/core/exporter"
)

func newDesc(subsystem string, name string, help string, labels prometheus.Labels, hiddenNamespace... bool) *prometheus.Desc {
	namespace := exporter.Namespace
	if len(hiddenNamespace) > 0 && hiddenNamespace[0] {
		namespace = ""
	}

	return prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, name),
		help, nil, labels,
	)
}

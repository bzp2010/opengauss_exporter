package scraper

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"opengauss_exporter/internal/core/scrape"
	"opengauss_exporter/internal/utils"
)

/**
  This scraper comes from GS_TOTAL_MEMORY_DETAIL
  view gets the memory used by the current node.
*/

type GsTotalMemoryDetailScraper struct {
}

func NewGsTotalMemoryDetailScraper() *GsTotalMemoryDetailScraper {
	return &GsTotalMemoryDetailScraper{}
}

type gsTotalMemoryDetail struct {
	nodename, memorytype string
	memorymbytes         int
}

func (s *gsTotalMemoryDetail) metric(labels prometheus.Labels) prometheus.Metric {
	var (
		name      = strings.ToLower(s.memorytype)
		subsystem = "total_memory_detail"
		val       = s.memorymbytes
	)

	labels["node_name"] = s.nodename

	desc := newDesc(subsystem, name, "", labels)

	return prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, float64(val))
}

func (g GsTotalMemoryDetailScraper) Scrape(t *scrape.Task) ([]prometheus.Metric, []error, error) {
	server := t.Fingerprint

	utils.GetLogger().Infow("Query GS_TOTAL_MEMORY_DETAIL view",
		"server", server,
	)

	query := "SELECT * FROM GS_TOTAL_MEMORY_DETAIL;"
	rows, err := t.DB.Query(query)
	if err != nil {
		return nil, nil, fmt.Errorf("Query GS_TOTAL_MEMORY_DETAIL view, server: %s, %v", server, err)
	}
	defer rows.Close()

	metrics := make([]prometheus.Metric, 0)
	for rows.Next() {
		totalMemoryDetail := &gsTotalMemoryDetail{}
		err = rows.Scan(&totalMemoryDetail.nodename, &totalMemoryDetail.memorytype, &totalMemoryDetail.memorymbytes)
		if err != nil {
			return nil, nil, fmt.Errorf("Query GS_TOTAL_MEMORY_DETAIL view failed, server: %s, %v", server, err)
		}

		if metric := totalMemoryDetail.metric(t.ConstLabels); metric != nil {
			metrics = append(metrics, metric)
		}
	}

	return metrics, nil, nil
}

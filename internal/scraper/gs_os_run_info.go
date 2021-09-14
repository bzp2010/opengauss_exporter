package scraper

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"opengauss_exporter/internal/core/scrape"
	"opengauss_exporter/internal/utils"
)

/**
  This scraper comes from GS_OS_RUN_INFO view
  obtains the operating system information of
  the current node.
*/

type GsOSRunInfoScraper struct {
}

func NewGsOSRunInfoScraper() *GsOSRunInfoScraper {
	return &GsOSRunInfoScraper{}
}

type gsOSRunInfo struct {
	name, comments string
	value          float64
	cumulative     bool
}

func (s *gsOSRunInfo) metric(labels prometheus.Labels) prometheus.Metric {
	var (
		name      = strings.ToLower(s.name)
		shortDesc = s.comments
		subsystem = "os_run_info"
		val       = s.value
	)

	desc := newDesc(subsystem, name, shortDesc, labels)

	switch s.name {
	case "NUM_CPUS", "NUM_CPU_CORES", "NUM_CPU_SOCKETS", "LOAD", "PHYSICAL_MEMORY_BYTES":
		return prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, val)
	case "IDLE_TIME", "BUSY_TIME", "USER_TIME", "SYS_TIME", "IOWAIT_TIME", "NICE_TIME":
		return prometheus.MustNewConstMetric(desc, prometheus.CounterValue, val)
	case "AVG_IDLE_TIME", "AVG_BUSY_TIME", "AVG_USER_TIME", "AVG_SYS_TIME", "AVG_IOWAIT_TIME", "AVG_NICE_TIME":
		return prometheus.MustNewConstMetric(desc, prometheus.CounterValue, val)
	case "VM_PAGE_IN_BYTES", "VM_PAGE_OUT_BYTES":
		return prometheus.MustNewConstMetric(desc, prometheus.CounterValue, val)
	}

	return nil
}

func (g GsOSRunInfoScraper) Scrape(t *scrape.Task) ([]prometheus.Metric, []error, error) {
	server := t.DataSource().Fingerprint()

	utils.GetLogger().Infow("Query GS_OS_RUN_INFO view",
		"server", server,
	)

	query := "SELECT name, value, comments, cumulative FROM GS_OS_RUN_INFO;"
	rows, err := t.DB().Query(query)
	if err != nil {
		return nil, nil, fmt.Errorf("Query GS_OS_RUN_INFO view, server: %s, %v", server, err)
	}
	defer rows.Close()

	metrics := make([]prometheus.Metric, 0)
	for rows.Next() {
		osRunInfo := &gsOSRunInfo{}
		err = rows.Scan(&osRunInfo.name, &osRunInfo.value, &osRunInfo.comments, &osRunInfo.cumulative)
		if err != nil {
			return nil, nil, fmt.Errorf("Query GS_OS_RUN_INFO view failed, server: %s, %v", server, err)
		}

		if metric := osRunInfo.metric(t.ConstLabels()); metric != nil {
			metrics = append(metrics, metric)
		}
	}

	return metrics, nil, nil
}

package scraper

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"opengauss_exporter/internal/core/exporter"
	"opengauss_exporter/internal/core/scrape"
	"opengauss_exporter/internal/utils"
)

/**
  This scraper comes from GS_INSTANCE_TIME view
  obtains the time information related to the
  current instance.
*/

type GsInstanceTimeScraper struct {
}

func NewGsInstanceTimeScraper() *GsInstanceTimeScraper {
	return &GsInstanceTimeScraper{}
}

type gsInstanceTime struct {
	stat_name string
	value     float64
}

func (s *gsInstanceTime) metric(labels prometheus.Labels) prometheus.Metric {
	var (
		name      = strings.ToLower(s.stat_name)
		subsystem = "instance_time"
		val       = s.value
	)

	desc := newDesc(subsystem, name, "", labels)

	return prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, val)
}

func (g GsInstanceTimeScraper) Scrape(t *scrape.Task) ([]prometheus.Metric, []error, error) {
	server := t.ConstLabels()[exporter.LabelServer]

	utils.GetLogger().Infof("Query GS_INSTANCE_TIME view: server: %s", server)

	query := "SELECT stat_name, value FROM GS_INSTANCE_TIME;"
	rows, err := t.DB().Query(query)
	if err != nil {
		return nil, nil, fmt.Errorf("Query GS_INSTANCE_TIME view, server: %s, %v", server, err)
	}
	defer rows.Close()

	metrics := make([]prometheus.Metric, 0)
	for rows.Next() {
		osRunInfo := &gsInstanceTime{}
		err = rows.Scan(&osRunInfo.stat_name, &osRunInfo.value)
		if err != nil {
			return nil, nil, fmt.Errorf("Query GS_INSTANCE_TIME view failed, server: %s, %v", server, err)
		}

		if metric := osRunInfo.metric(t.ConstLabels()); metric != nil {
			metrics = append(metrics, metric)
		}
	}

	return metrics, nil, nil
}

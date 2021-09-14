package scraper

import (
	"fmt"
	"reflect"

	"github.com/prometheus/client_golang/prometheus"

	"opengauss_exporter/internal/core/scrape"
	"opengauss_exporter/internal/utils"
)

/**
  This scraper comes from GS_SQL_COUNT view obtains
  SQL metering data distinguished by database node
  name and login username.
*/

type GsSQLCountScraper struct {
}

func NewGsSQLCountScraper() *GsSQLCountScraper {
	return &GsSQLCountScraper{}
}

type gsSQLCount struct {
	node_name, user_name                   string
	select_count, update_count             int64
	insert_count, delete_count             int64
	mergeinto_count, ddl_count             int64
	dml_count, dcl_count                   int64
	total_select_elapse, avg_select_elapse int64
	max_select_elapse, min_select_elapse   int64
	total_update_elapse, avg_update_elapse int64
	max_update_elapse, min_update_elapse   int64
	total_insert_elapse, avg_insert_elapse int64
	max_insert_elapse, min_insert_elapse   int64
	total_delete_elapse, avg_delete_elapse int64
	max_delete_elapse, min_delete_elapse   int64
}

func (s *gsSQLCount) metric(labels prometheus.Labels) []prometheus.Metric {
	var (
		subsystem = "sql_count"
		metrics   = make([]prometheus.Metric, 0)
	)

	labels["node_name"] = s.node_name
	labels["user_name"] = s.user_name

	t := reflect.TypeOf(*s)
	v := reflect.ValueOf(*s)
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Name
		if key == "node_name" || key == "user_name" {
			continue
		}
		val := float64(v.FieldByName(key).Int())

		desc := newDesc(subsystem, key, "", labels)
		switch key {
		case "select_count", "update_count", "insert_count", "delete_count", "mergeinto_count", "ddl_count", "dml_count", "dcl_count":
			metrics = append(metrics, prometheus.MustNewConstMetric(desc, prometheus.CounterValue, val))
		case "total_select_elapse", "total_update_elapse", "total_insert_elapse", "total_delete_elapse":
			metrics = append(metrics, prometheus.MustNewConstMetric(desc, prometheus.CounterValue, val))
		case "avg_select_elapse", "avg_update_elapse", "avg_insert_elapse", "avg_delete_elapse":
			metrics = append(metrics, prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, val))
		case "max_select_elapse", "max_update_elapse", "max_insert_elapse", "max_delete_elapse":
			metrics = append(metrics, prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, val))
		case "min_select_elapse", "min_update_elapse", "min_insert_elapse", "min_delete_elapse":
			metrics = append(metrics, prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, val))
		}
	}

	return metrics
}

func (g GsSQLCountScraper) Scrape(t *scrape.Task) ([]prometheus.Metric, []error, error) {
	server := t.DataSource().Fingerprint()

	utils.GetLogger().Infow("Query GS_SQL_COUNT view",
		"server", server,
	)

	query := "SELECT * FROM GS_SQL_COUNT;"
	rows, err := t.DB().Query(query)
	if err != nil {
		return nil, nil, fmt.Errorf("Query GS_SQL_COUNT view, server: %s, %v", server, err)
	}
	defer rows.Close()

	metrics := make([]prometheus.Metric, 0)
	for rows.Next() {
		sqlCount := &gsSQLCount{}
		err = rows.Scan(
			&sqlCount.node_name, &sqlCount.user_name,
			&sqlCount.select_count, &sqlCount.update_count,
			&sqlCount.insert_count, &sqlCount.delete_count,
			&sqlCount.mergeinto_count, &sqlCount.ddl_count,
			&sqlCount.dml_count, &sqlCount.dcl_count,
			&sqlCount.total_select_elapse, &sqlCount.avg_select_elapse,
			&sqlCount.max_select_elapse, &sqlCount.min_select_elapse,
			&sqlCount.total_update_elapse, &sqlCount.avg_update_elapse,
			&sqlCount.max_update_elapse, &sqlCount.avg_update_elapse,
			&sqlCount.total_insert_elapse, &sqlCount.avg_insert_elapse,
			&sqlCount.max_insert_elapse, &sqlCount.min_insert_elapse,
			&sqlCount.total_delete_elapse, &sqlCount.avg_delete_elapse,
			&sqlCount.max_delete_elapse, &sqlCount.min_delete_elapse,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("Query GS_SQL_COUNT view failed, server: %s, %v", server, err)
		}

		metrics = append(metrics, sqlCount.metric(t.ConstLabels())...)
	}

	return metrics, nil, nil
}

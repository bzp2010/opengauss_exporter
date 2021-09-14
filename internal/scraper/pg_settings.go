package scraper

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"opengauss_exporter/internal/core/scrape"
	"opengauss_exporter/internal/utils"
)

/**
  This scraper is from PG_SETTINGS view
  obtains the configuration information
  of the current node.
*/

type PgSettingsScraper struct {
}

func NewPgSettingsScraper() *PgSettingsScraper {
	return &PgSettingsScraper{}
}

// pgSetting is represents a PostgreSQL runtime variable as returned by the
// pg_settings view.
// source from postgres_exporter
type pgSetting struct {
	name, setting, unit, shortDesc, vartype string
}

// source from postgres_exporter
func (s *pgSetting) metric(labels prometheus.Labels) prometheus.Metric {
	var (
		err       error
		name      = strings.Replace(s.name, ".", "_", -1)
		unit      = s.unit // nolint: ineffassign
		shortDesc = s.shortDesc
		subsystem = "pg_settings"
		val       float64
	)

	switch s.vartype {
	case "bool":
		if s.setting == "on" {
			val = 1
		}
	case "integer", "real":
		if val, unit, err = s.normaliseUnit(); err != nil {
			// Panic, since we should recognise all units
			// and don't want to silently exlude metrics
			panic(err)
		}

		if len(unit) > 0 {
			name = fmt.Sprintf("%s_%s", name, unit)
			shortDesc = fmt.Sprintf("%s [Units converted to %s.]", shortDesc, unit)
		}
	default:
		// Panic because we got a type we didn't ask for
		panic(fmt.Sprintf("Unsupported vartype %q", s.vartype))
	}

	desc := newDesc(subsystem, name, shortDesc, labels, true)
	return prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, val)
}

// source from postgres_exporter
func (s *pgSetting) normaliseUnit() (val float64, unit string, err error) {
	val, err = strconv.ParseFloat(s.setting, 64)
	if err != nil {
		return val, unit, fmt.Errorf("Error converting setting %q value %q to float: %s", s.name, s.setting, err)
	}

	// Units defined in: https://www.postgresql.org/docs/current/static/config-setting.html
	switch s.unit {
	case "":
		return
	case "ms", "s", "min", "h", "d":
		unit = "seconds"
	case "B", "kB", "MB", "GB", "TB", "8kB", "16kB", "32kB", "16MB", "32MB", "64MB":
		unit = "bytes"
	default:
		err = fmt.Errorf("Unknown unit for runtime variable: %q", s.unit)
		return
	}

	// -1 is special, don't modify the value
	if val == -1 {
		return
	}

	switch s.unit {
	case "ms":
		val /= 1000
	case "min":
		val *= 60
	case "h":
		val *= 60 * 60
	case "d":
		val *= 60 * 60 * 24
	case "kB":
		val *= math.Pow(2, 10)
	case "MB":
		val *= math.Pow(2, 20)
	case "GB":
		val *= math.Pow(2, 30)
	case "TB":
		val *= math.Pow(2, 40)
	case "8kB":
		val *= math.Pow(2, 13)
	case "16kB":
		val *= math.Pow(2, 14)
	case "32kB":
		val *= math.Pow(2, 15)
	case "16MB":
		val *= math.Pow(2, 24)
	case "32MB":
		val *= math.Pow(2, 25)
	case "64MB":
		val *= math.Pow(2, 26)
	}

	return
}

func (p PgSettingsScraper) Scrape(t *scrape.Task) ([]prometheus.Metric, []error, error) {
	server := t.DataSource().Fingerprint()

	utils.GetLogger().Infow("Query PG_SETTINGS view",
		"server", server,
	)

	query := "SELECT name, setting, COALESCE(unit, ''), short_desc, vartype FROM pg_settings WHERE vartype IN ('bool', 'integer', 'real');"
	rows, err := t.DB().Query(query)
	if err != nil {
		return nil, nil, fmt.Errorf("Query PG_SETTINGS view failed, server: %s, %v", server, err)
	}
	defer rows.Close()

	metrics := make([]prometheus.Metric, 0)
	for rows.Next() {
		settings := &pgSetting{}
		err = rows.Scan(&settings.name, &settings.setting, &settings.unit, &settings.shortDesc, &settings.vartype)
		if err != nil {
			return nil, nil, fmt.Errorf("Query PG_SETTINGS view failed, server: %s, %v", server, err)
		}

		metrics = append(metrics, settings.metric(t.ConstLabels()))
	}

	return metrics, nil, nil
}

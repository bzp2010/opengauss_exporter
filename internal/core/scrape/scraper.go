package scrape

import "github.com/prometheus/client_golang/prometheus"

type Scraper interface {
	// Scrape metrics
	// first []error is non fatal errors
	// second error is fatal error
	Scrape(task *Task) ([]prometheus.Metric, []error, error)
}

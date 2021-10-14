package scraper

import (
	"github.com/prometheus/client_golang/prometheus"

	"opengauss_exporter/internal/core/scrape"
)

type ExampleScraper struct {}

func NewExampleScraper() *ExampleScraper {
	return &ExampleScraper{}
}

func (e ExampleScraper) Scrape(task *scrape.Task) ([]prometheus.Metric, []error, error) {
	// execute SQL in database and return fetched data
	// database connection in `task.DB`

	return []prometheus.Metric{}, nil, nil
}



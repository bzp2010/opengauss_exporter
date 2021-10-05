package scrape

import (
	"database/sql"
	"fmt"
	"regexp"
	"time"

	_ "gitee.com/opengauss/openGauss-connector-go-pq"
	"github.com/blang/semver"
	"github.com/gogf/gf/os/gtimer"
	"github.com/prometheus/client_golang/prometheus"

	"opengauss_exporter/internal/config"
	"opengauss_exporter/internal/core/cache"
	"opengauss_exporter/internal/core/exporter"
	"opengauss_exporter/internal/utils"
)

type Option func(*Task)
type options struct{}

type Task struct {
	options         *options
	timer           *gtimer.Timer
	dataSource      config.DataSource
	db              *sql.DB
	pgVersion       semver.Version
	ogVersion       semver.Version
	connectMaxRetry int
	connectRetry    int

	constLabels  prometheus.Labels
	duration     prometheus.Gauge
	error        prometheus.Gauge
	dataSourceUp prometheus.Gauge
	version      prometheus.Gauge
	totalScrapes prometheus.Counter

	scrapers []Scraper

	errSig  chan error
	stopSig chan bool
}

// NewTask create a database scrape task
func NewTask(opts ...Option) *Task {
	t := &Task{
		options: &options{},
	}

	// setup options
	for _, opt := range opts {
		opt(t)
	}

	// setup timer
	if t.timer == nil {
		t.timer = gtimer.New()
	}

	// setup labels
	t.constLabels = prometheus.Labels{
		exporter.LabelServer: t.dataSource.Fingerprint(),
	}

	// setup connect tries
	t.connectRetry = 1
	t.connectMaxRetry = t.dataSource.MaxRetry

	t.pgVersion = semver.MustParse("9.2.4")
	t.ogVersion = semver.MustParse("2.0.1")

	return t
}

func (t *Task) With(opt Option) *Task {
	opt(t)
	return t
}

func (t *Task) DB() *sql.DB {
	return t.db
}

func (t *Task) DataSource() config.DataSource {
	return t.dataSource
}

func (t *Task) ConstLabels() prometheus.Labels {
	return t.constLabels
}

func (t *Task) PostgreSQLVersion() semver.Version {
	return t.pgVersion
}

func (t *Task) OpenGaussVersion() semver.Version {
	return t.ogVersion
}

func (t *Task) Start() {
	// initialize Task metrics
	t.setupTaskMetrics()

	// retry logic for connect to datasource
RETRY:
	needRetry, err := t.connectDataSource()
	if err != nil {
		// retry
		if needRetry && t.connectRetry < t.connectMaxRetry {
			t.connectRetry++
			goto RETRY
		}

		t.errSig <- fmt.Errorf("connect to datasource failed after max %d retry: server: %s err: %v", t.connectRetry, t.dataSource.Fingerprint(), err)
		return
	}

	// setup scrape timer
	t.timer.AddSingleton(t.dataSource.Duration, func() {
		t.scrape()
	})
	t.timer.Start()
}

func (t *Task) Stop() {
	t.timer.Stop()
	if t.db != nil {
		_ = t.db.Close()
	}
	return
}

func (t *Task) connectDataSource() (bool, error) {
	// check datasource config
	utils.GetLogger().Infof("Connect datasource #%d: server: %s", t.connectRetry, t.dataSource.Fingerprint())
	db, err := sql.Open("opengauss", t.dataSource.DSN)
	if err != nil {
		// config error: not need retry
		return false, fmt.Errorf("Check datasource failed: server: %s err: %v"+t.dataSource.Fingerprint(), err)
	}

	// connect to datasource and check
	err = db.Ping()
	if err != nil { // connect failed
		// connect error: need retry
		return true, fmt.Errorf("Connect datasource failed #%d: server: %s err: %v",
			t.connectRetry, t.dataSource.Fingerprint(), err,
		)
	}

	utils.GetLogger().Infow("Connect datasource successful",
		"server", t.dataSource.Fingerprint(),
	)

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	t.db = db
	t.dataSourceUp.Set(1)

	// fetch database version
	err = t.queryDBVersion()
	if err != nil {
		return false, err
	}

	return false, nil
}

func (t *Task) queryDBVersion() error {
	rows, err := t.db.Query("SELECT * FROM version();")
	if err != nil {
		return fmt.Errorf("Database version query failed: %v", err)
	}
	defer rows.Close()

	rows.Next()

	var versionStr string
	err = rows.Scan(&versionStr)
	if err != nil {
		return fmt.Errorf("Database version scan failed: %v", err)
	}

	submatches := regexp.MustCompile(`^\w+ ((\d+)(\.\d+)?(\.\d+)?) \(\w+ ((\d+)(\.\d+)?(\.\d+)?) \w+ \w+\)`).FindStringSubmatch(versionStr)
	if len(submatches) > 1 {
		t.pgVersion, err = semver.ParseTolerant(submatches[1])
		if err != nil {
			return fmt.Errorf("Database PostgreSQL version parse failed: %v", err)
		}

		t.ogVersion, err = semver.ParseTolerant(submatches[5])
		if err != nil {
			return fmt.Errorf("Database openGauss version parse failed: %v", err)
		}
	}

	// generate version labels
	versionInfo := make(map[string]string)
	for k, v := range t.constLabels {
		versionInfo[k] = v
	}
	versionInfo["pg"] = t.pgVersion.String()
	versionInfo["og"] = t.ogVersion.String()

	// set version information
	t.version = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   exporter.NamespaceRoot,
		Name:        "version",
		Help:        "Version information of database nodes",
		ConstLabels: versionInfo,
	})

	return nil
}

func (t *Task) setupTaskMetrics() {
	t.duration = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   exporter.NamespaceRoot,
		Subsystem:   exporter.SubsystemScrapeTask,
		Name:        "last_scrape_duration_seconds",
		Help:        "Duration of the last scrape of metrics from openGauss.",
		ConstLabels: t.constLabels,
	})
	t.totalScrapes = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace:   exporter.NamespaceRoot,
		Subsystem:   exporter.SubsystemScrapeTask,
		Name:        "scrapes_total",
		Help:        "Total number of times openGauss was scraped for metrics.",
		ConstLabels: t.constLabels,
	})
	t.error = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   exporter.NamespaceRoot,
		Subsystem:   exporter.SubsystemScrapeTask,
		Name:        "last_scrape_error",
		Help:        "Whether the last scrape of metrics from openGauss resulted in an error (1 for error, 0 for success).",
		ConstLabels: t.constLabels,
	})
	t.dataSourceUp = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   exporter.NamespaceRoot,
		Name:        "up",
		Help:        "Whether the last scrape of metrics from openGauss was able to connect to the server (1 for yes, 0 for no).",
		ConstLabels: t.constLabels,
	})
}

// Scrape datasource metrics (triggered by Task timer)
func (t *Task) scrape() {
	// calculate used time
	defer func(begun time.Time) {
		t.duration.Set(time.Since(begun).Seconds())
	}(time.Now())

	metrics := make([]prometheus.Metric, 0)

	// run all scrapers
	for _, s := range t.scrapers {
		scraperMetrics, nonfatalErrors, err := s.Scrape(t)

		// handle fatal error
		if err != nil {
			t.error.Inc()
			t.errSig <- err
			utils.GetLogger().Errorf("Scrape task returned fatal errors: server: %s %v", t.dataSource.Fingerprint(), err)
			return
		}

		// handle non fatal error
		if nonfatalErrors != nil && len(nonfatalErrors) > 0 {
			for _, err := range nonfatalErrors {
				t.error.Inc()
				t.errSig <- err
			}

			utils.GetLogger().Warnf("Scrape task returned %d non fatal errors: server: %s", len(nonfatalErrors), t.dataSource.Fingerprint())
		}

		metrics = append(metrics, scraperMetrics...)
	}

	t.totalScrapes.Inc()

	// cache metrics
	_ = cache.Metrics.Set(t.dataSource.Fingerprint(), t.mergeTaskMetrics(metrics), 0)
}

func (t *Task) mergeTaskMetrics(scrapeMetrics []prometheus.Metric) []prometheus.Metric {
	taskMetrics := []prometheus.Metric{
		t.duration, t.totalScrapes, t.error, t.dataSourceUp, t.version,
	}

	return append(taskMetrics, scrapeMetrics...)
}

// WithDataSource set a datasource to Task
func WithDataSource(d config.DataSource) Option {
	return func(t *Task) {
		t.dataSource = d
	}
}

// WithScraper add a scraper to Task
func WithScraper(s Scraper) Option {
	return func(t *Task) {
		t.scrapers = append(t.scrapers, s)
	}
}

// WithErrorSig set a error channel to Task
func WithErrorSig(errSig chan error) Option {
	return func(t *Task) {
		t.errSig = errSig
	}
}

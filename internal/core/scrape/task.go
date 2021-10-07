package scrape

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	"time"

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
	config      config.Task
	options     *options
	name        string
	Fingerprint string

	PGVersion   semver.Version
	OGVersion   semver.Version
	ConstLabels prometheus.Labels

	// built-in metrics
	taskTotalScrapes prometheus.Counter
	taskDuration     prometheus.Gauge
	taskError        prometheus.Gauge
	taskName         prometheus.Gauge
	coreUp           prometheus.Gauge
	coreVersion      prometheus.Gauge

	// database connection
	DB           *sql.DB
	connectRetry int
	isConnected  bool

	// scrape timer
	timer *gtimer.Timer

	// task scraper list
	scrapers []Scraper

	// signal channel
	errSig  chan error
	stopSig chan bool
}

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
	t.ConstLabels = prometheus.Labels{
		exporter.LabelServer: t.Fingerprint,
	}

	t.PGVersion = semver.MustParse("9.2.4")

	return t
}

func (t *Task) Start() {
	// initialize built-in metrics
	t.setupTaskMetrics()

	// retry logic for connect to datasource
RETRY:
	needRetry, err := t.connectDataSource()
	if err != nil {
		t.isConnected = false
		t.scrape()

		// retry
		if needRetry {
			t.connectRetry++
			goto RETRY
		} else {
			t.errSig <- fmt.Errorf("Connect to database failed: server: %s err: %v", t.Fingerprint, err)
		}
		return
	}

	t.isConnected = true

	// setup scrape timer
	t.timer.AddSingleton(t.config.Duration, func() {
		t.scrape()
	})
	t.timer.Start()
}

func (t *Task) Stop() {
	t.timer.Stop()
	if t.DB != nil {
		_ = t.DB.Close()
	}
	return
}

func (t *Task) Config() config.Task {
	return t.config
}

func (t *Task) setupTaskMetrics() {
	t.taskDuration = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   exporter.NamespaceRoot,
		Subsystem:   exporter.SubsystemScrapeTask,
		Name:        "last_scrape_duration_seconds",
		Help:        "Duration of the last scrape of metrics from openGauss",
		ConstLabels: t.ConstLabels,
	})
	t.taskTotalScrapes = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace:   exporter.NamespaceRoot,
		Subsystem:   exporter.SubsystemScrapeTask,
		Name:        "scrapes_total",
		Help:        "Total number of times openGauss was scraped for metrics.",
		ConstLabels: t.ConstLabels,
	})
	t.taskError = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   exporter.NamespaceRoot,
		Subsystem:   exporter.SubsystemScrapeTask,
		Name:        "last_scrape_error",
		Help:        "Whether the last scrape of metrics from openGauss resulted in an error (1 for error, 0 for success).",
		ConstLabels: t.ConstLabels,
	})
	t.coreUp = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   exporter.NamespaceRoot,
		Name:        "up",
		Help:        "Whether the last scrape of metrics from openGauss was able to connect to the server (1 for yes, 0 for no).",
		ConstLabels: t.ConstLabels,
	})

	taskName := make(map[string]string)
	for k, v := range t.ConstLabels {
		taskName[k] = v
	}
	taskName["name"] = t.config.Name
	t.taskName = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   exporter.NamespaceRoot,
		Subsystem:   exporter.SubsystemScrapeTask,
		Name:        "name",
		Help:        "Scrape task name.",
		ConstLabels: taskName,
	})
}

func (t *Task) connectDataSource() (bool, error) {
	// check datasource config
	db, err := sql.Open("opengauss", t.config.DSN)
	if err != nil {
		// config error: not need retry
		utils.GetLogger().Errorw("DSN config error",
			"server", t.Fingerprint,
			"error", err,
		)
		return false, fmt.Errorf("Check datasource failed: %v", err)
	}

	// set database connection limit
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	// connect to datasource and check
	utils.GetLogger().Infow("Connecting to database",
		"server", t.Fingerprint,
		"retry", t.connectRetry,
	)
	err = db.Ping()
	if err != nil { // connect failed
		// connect error: need retry
		utils.GetLogger().Errorw("Connect to database failed",
			"server", t.Fingerprint,
			"error", err,
		)
		return true, fmt.Errorf("Connect database failed: %v", err)
	}
	utils.GetLogger().Infow("Connect database successful",
		"server", t.Fingerprint,
	)

	t.DB = db

	// fetch database version
	err = t.queryDBVersion()
	if err != nil {
		return false, err
	}

	return false, nil
}

func (t *Task) queryDBVersion() error {
	rows, err := t.DB.Query("SELECT version();")
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

	subMatches := regexp.MustCompile(`openGauss (\d+)(\.\d+)?(\.\d+)?`).FindStringSubmatch(versionStr)
	if len(subMatches) > 1 {
		t.OGVersion, err = semver.ParseTolerant(strings.Replace(subMatches[0], "openGauss", "", 1))
		if err != nil {
			return fmt.Errorf("Database openGauss version parse failed: %v", err)
		}
	}

	// generate version labels
	versionInfo := make(map[string]string)
	for k, v := range t.ConstLabels {
		versionInfo[k] = v
	}
	versionInfo["pg"] = t.PGVersion.String()
	versionInfo["og"] = t.OGVersion.String()

	// set version information
	t.coreVersion = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   exporter.NamespaceRoot,
		Name:        "version",
		Help:        "Version information of database nodes",
		ConstLabels: versionInfo,
	})

	return nil
}

// Scrape datasource metrics (triggered by Task timer)
func (t *Task) scrape() {
	// calculate used time
	defer func(begun time.Time) {
		t.taskDuration.Set(time.Since(begun).Seconds())
	}(time.Now())
	t.taskError.Set(0)

	if !t.isConnected {
		t.coreUp.Set(0)
		_ = cache.Metrics.Set(t.Fingerprint, []prometheus.Metric{t.taskName, t.coreUp}, 0)
		return
	}
	t.coreUp.Set(1)

	metrics := make([]prometheus.Metric, 0)

	// run all scrapers
	for _, s := range t.scrapers {
		scraperMetrics, nonfatalErrors, err := s.Scrape(t)

		// handle fatal error
		if err != nil {
			t.taskError.Set(1)
			t.errSig <- err
			utils.GetLogger().Errorf("Scrape task returned fatal errors: server: %s %v", t.Fingerprint, err)
			return
		}

		// handle non fatal error
		if nonfatalErrors != nil && len(nonfatalErrors) > 0 {
			for _, err := range nonfatalErrors {
				t.taskError.Set(1)
				t.errSig <- err
			}

			utils.GetLogger().Warnf("Scrape task returned %d non fatal errors: server: %s", len(nonfatalErrors), t.Fingerprint)
		}

		metrics = append(metrics, scraperMetrics...)
	}

	t.taskTotalScrapes.Inc()

	// cache metrics
	_ = cache.Metrics.Set(t.Fingerprint, t.mergeTaskMetrics(metrics), 0)
}

func (t *Task) mergeTaskMetrics(scrapeMetrics []prometheus.Metric) []prometheus.Metric {
	taskMetrics := []prometheus.Metric{
		t.taskDuration, t.taskTotalScrapes, t.taskError, t.taskName,
		t.coreUp, t.coreVersion,
	}

	return append(taskMetrics, scrapeMetrics...)
}

// WithConfig set a task config
func WithConfig(taskConfig config.Task) Option {
	return func(t *Task) {
		t.config = taskConfig
		t.Fingerprint = t.config.Fingerprint()
		if t.config.Name != "" {
			t.name = t.config.Name
		} else {
			t.name = t.Fingerprint
		}

		// scrapers
		for _, scraperId := range t.config.Scrapers {
			scraper, ok := scraperList[scraperId]
			if ok {
				t.scrapers = append(t.scrapers, scraper)
			} else {
				utils.GetLogger().Warnw("Task use unsupported scraper",
					"server", t.Fingerprint,
					"id", scraperId,
				)
			}
		}
	}
}

// WithErrorSig set a error channel to Task
func WithErrorSig(errSig chan error) Option {
	return func(t *Task) {
		t.errSig = errSig
	}
}

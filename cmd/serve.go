package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"opengauss_exporter/internal/config"
	"opengauss_exporter/internal/core/exporter"
	"opengauss_exporter/internal/core/scrape"
	"opengauss_exporter/internal/core/server"
	"opengauss_exporter/internal/scraper"
	"opengauss_exporter/internal/utils"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run openGauss Exporter",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runExporter()
	},
}

func runExporter() error {
	serverErrSig := make(chan error, 100)

	// load and parse config
	err := config.Init(configFile)
	if err != nil {
		return err
	}

	// setup mainly modules
	setupPrometheusCollector()
	scrapeManager := setupScrapeManager()
	serverManager := setupServerManager(serverErrSig)

	// graceful exit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrSig:
		panic(fmt.Errorf("HTTP server start error: %v", err))
		return err
	case <-quit:
		scrapeManager.Stop()
		serverManager.Stop()
		utils.GetLogger().Infof("See you next time!")
	}

	return nil
}

func setupPrometheusCollector() {
	prometheus.MustRegister(version.NewCollector(exporter.Name))
	prometheus.MustRegister(exporter.NewExporter())
}

func setupServerManager(errSig chan error) *server.Manager {
	serverManager := server.NewManager(
		server.WithHTTPServer(
			viper.GetString("server.http.host"),
			viper.GetString("server.http.port"),
		),
	)
	serverManager.Start(errSig)

	return serverManager
}

func setupScrapeManager() *scrape.Manager {
	scrapeErrSig := make(chan error, 10000)
	scrapeManager := scrape.NewManager()
	for _, source := range config.DataSources {
		task := scrape.NewTask(
			scrape.WithDataSource(source),
			scrape.WithErrorSig(scrapeErrSig),
		)

		if source.EnableSettings {
			task.With(scrape.WithScraper(scraper.NewPgSettingsScraper()))
		}

		if source.EnableOSRunInfo {
			task.With(scrape.WithScraper(scraper.NewGsOSRunInfoScraper()))
		}

		if source.EnableTotalMemoryDetail {
			task.With(scrape.WithScraper(scraper.NewGsTotalMemoryDetailScraper()))
		}

		if source.EnableSQLCount {
			task.With(scrape.WithScraper(scraper.NewGsSQLCountScraper()))
		}

		if source.EnableInstanceTime {
			task.With(scrape.WithScraper(scraper.NewGsInstanceTimeScraper()))
		}

		if source.EnablePostgreSQLExporter {
			task.With(scrape.WithScraper(scraper.NewBuiltinSQLScraper()))
		}

		scrapeManager.AddTask(task)
	}
	scrapeManager.Start()

	// scrape error handler
	go func() {
		for err := range scrapeErrSig {
			utils.GetLogger().Errorf("Scrape task error: %s", err.Error())
		}
	}()

	return scrapeManager
}

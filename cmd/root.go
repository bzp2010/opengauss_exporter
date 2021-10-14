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

var configFile string

var rootCmd = &cobra.Command{
	Use:   "opengauss_exporter",
	Short: "OpenGauss metrics exporter",
	Long:  `A OpenGauss metric exporter for Prometheus`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runExporter()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yaml", "config file path")
}

func runExporter() error {
	serverErrSig := make(chan error, 100)

	// load and parse config
	err := config.Init(configFile)
	if err != nil {
		return err
	}

	// init scrapers
	scraper.Init()

	// setup mainly modules
	setupPrometheusCollector()
	scrapeManager := setupScrapeManager()
	serverManager := setupServerManager(serverErrSig)

	// graceful exit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrSig:
		panic(fmt.Errorf("HTTP(s) server start error: %v", err))
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
		server.WithMiddlewares(viper.GetStringMap("server.middlewares")),
	)

	if viper.IsSet("server.https") {
		serverManager.With(server.WithHTTPSServer(
			viper.GetString("server.https.host"),
			viper.GetString("server.https.port"),
			viper.GetString("server.https.cert"),
			viper.GetString("server.https.key"),
		))
	}

	serverManager.Start(errSig)

	return serverManager
}

func setupScrapeManager() *scrape.Manager {
	scrapeErrSig := make(chan error, 10000)
	scrapeManager := scrape.NewManager()
	for _, taskConfig := range config.Tasks {
		task := scrape.NewTask(
			scrape.WithConfig(taskConfig),
			scrape.WithErrorSig(scrapeErrSig),
		)
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

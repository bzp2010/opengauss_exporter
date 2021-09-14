package config

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/gogf/gf/util/gconv"
	"github.com/spf13/viper"

	"opengauss_exporter/internal/utils"
)

type DataSource struct {
	DSN      string
	Duration time.Duration
	MaxRetry int
	Master   bool

	// features
	EnableSettings           bool
	EnableOSRunInfo          bool `c:"enable_os_run_info"`
	EnableTotalMemoryDetail  bool
	EnableSQLCount           bool `c:"enable_sql_count"`
	EnableInstanceTime       bool
	EnablePostgreSQLExporter bool `c:"enable_postgresql_exporter"`
}

func (d DataSource) Fingerprint() string {
	serverURL, _ := url.Parse(d.DSN)
	return serverURL.Host
}

var (
	DataSources = make([]DataSource, 0)
)

func Init(configFile string) error {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	err := viper.ReadInConfig()
	if err != nil {
		utils.GetLogger().Errorf("Config file load failed: %v", err)
		return err
	}
	utils.GetLogger().Infow("Config file load successful",
		"path", viper.ConfigFileUsed(),
	)

	err = setupServer()
	if err != nil {
		return err
	}

	err = setupDataSources()
	if err != nil {
		return err
	}

	return nil
}

func setupServer() error {
	return nil
}

func setupDataSources() error {
	for key, val := range viper.Get("data_sources").([]interface{}) {
		var dataSource DataSource
		err := gconv.Struct(val, &dataSource)
		if err != nil {
			return errors.New(fmt.Sprintf("Data source parse failed: index: %d", key))
		}
		DataSources = append(DataSources, dataSource)
	}
	return nil
}

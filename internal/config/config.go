package config

import (
	"net/url"
	"time"

	"github.com/spf13/viper"

	"opengauss_exporter/internal/utils"
)

type Task struct {
	DSN      string
	Name     string
	Duration time.Duration
	Master   bool

	Scrapers []string
}

func (t Task) Fingerprint() string {
	serverURL, _ := url.Parse(t.DSN)
	return serverURL.Host
}

var (
	Tasks = make([]Task, 0)
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

	err = setupTasks()
	if err != nil {
		return err
	}

	return nil
}

func setupServer() error {
	return nil
}

func setupTasks() error {
	err := viper.UnmarshalKey("tasks", &Tasks)
	if err != nil {
		return err
	}
	return nil
}

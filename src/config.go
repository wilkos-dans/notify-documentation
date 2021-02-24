package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Debugging              bool   `yaml:"debugging"`
	WebrootFolderPath      string `yaml:"webroot"`
	ScenariosCsvDataSource string `yaml:"scenarios_csv_source"`
	NotificationsCsvDataSource string `yaml:"notifications_csv_source"`
	NotificationPatternsCsvSource string `yaml:"notification_patterns_csv_source"`
}

func (config *Config) unmarshal(filePath string) error {
	configData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(configData), config)
	if err != nil {
		return err
	}
	if config.Debugging == true {
		zapLogger, _ = configureZapLogger(true)
		zapLogger.Info("Debugging enabled")
	}
	zapLogger.Info(fmt.Sprintf("Webroot folder path set to %s", config.WebrootFolderPath))
	return err
}

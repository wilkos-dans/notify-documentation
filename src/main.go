package main

import (
	"flag"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger
var config = Config{}
var website = Website{}

//var Scenarios []Scenario

func main() {
	var err error
	var configFilePath string
	zapLogger, _ = configureZapLogger(false)
	//	### Load arguments
	flag.StringVar(&configFilePath, "c", "", "path to a valid config yaml file")
	flag.Parse()

	// ### Load configuration
	err = (&config).unmarshal(configFilePath)
	if err != nil {
		zapLogger.Fatal("Unable to initialise - halting execution")
	} else {
		zapLogger.Info("Configuration loaded OK")
	}
	err = website.Initialise(config.WebrootFolderPath, config.ScenariosCsvDataSource, config.NotificationsCsvDataSource, config.PatternsCsvDataSource)
	if err != nil {
		zapLogger.Fatal("Unable to initialise - halting execution")
	} else {
		zapLogger.Info("Website initialised OK")
	}

	//scenario1 := website.GetScenarioById("repository_requests_review_OLD")
	//for _,notification := range scenario1.Notifications {
	//	fmt.Printf("WorkflowStep: %s \n",notification.Step.From)
	//}

	website.WriteWebPages()
	//website.WriteNotificationWebPages()
}

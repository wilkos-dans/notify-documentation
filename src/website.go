package main

import (
	"fmt"
	"github.com/jszwec/csvutil"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Website struct {
	WebrootFolderPath               string
	ContentFolderPath               string
	ScenariosFolderPath             string
	NotificationsFolderPath         string
	NotificationPatternsFolderPath string
	StaticContentFolderPath         string
	Scenarios                       []*Scenario
	NotificationsPatterns          []NotificationPattern
}

func (website *Website) Initialise(webrootPath, scenariosCsvSource, notificationsCsvSource, notificationPatternsCsvSource string) error {
	var err error
	website.WebrootFolderPath = webrootPath
	website.ContentFolderPath = filepath.Join(website.WebrootFolderPath, "content")
	website.ScenariosFolderPath = filepath.Join(website.ContentFolderPath, "scenarios")
	website.NotificationsFolderPath = filepath.Join(website.ContentFolderPath, "notifications")
	website.NotificationPatternsFolderPath = filepath.Join(website.ContentFolderPath, "notification_patterns")
	website.StaticContentFolderPath = filepath.Join(website.WebrootFolderPath, "static")
	scenariosCsvInput, err := readBytesFromFileOrUrl(scenariosCsvSource)
	if err != nil {
		return err
	}
	err = csvutil.Unmarshal(scenariosCsvInput, &website.Scenarios)
	if err != nil {
		return err
	}
	zapLogger.Info("Scenarios loaded OK")
	notificationsCsvInput, err := readBytesFromFileOrUrl(notificationsCsvSource)
	if err != nil {
		zapLogger.Error(err.Error())
		return err
	}
	var notificationConfigs []NotificationConfig
	err = csvutil.Unmarshal(notificationsCsvInput, &notificationConfigs)
	if err != nil {
		zapLogger.Error(err.Error())
		return err
	}
	for _, notificationConfig := range notificationConfigs {
		scenario := website.GetScenarioById(notificationConfig.ScenarioId)
		if scenario != nil {
			notification := Notification{
				ScenarioId:  notificationConfig.ScenarioId,
				Title:       notificationConfig.Title,
				Description: notificationConfig.Description,
				Updated:     time.Now().Format("2006-01-02"),
				PatternId: notificationConfig.PatternId,
				Step: StepStruct{
					Scope:    notificationConfig.Scope,
					Position: notificationConfig.Position,
					Sender:   notificationConfig.Sender,
				},
			}
			scenario.Notifications = append(scenario.Notifications, notification)
		}
	}
	notificationPatternsCsvInput, err := readBytesFromFileOrUrl(notificationPatternsCsvSource)
	if err != nil {
		zapLogger.Error(err.Error())
		return err
	}
	err = csvutil.Unmarshal(notificationPatternsCsvInput, &website.NotificationsPatterns)
	if err != nil {
		zapLogger.Error(err.Error())
		return err
	}
	zapLogger.Info("Notification Patterns loaded OK")
	return err
}

func (website *Website) GetScenarioById(id string) *Scenario {
	for _, scenario := range website.Scenarios {
		if scenario.Id == id {
			return scenario
		}
	}
	return nil
}

func (website *Website) WriteWebPages() error {
	var err error
	for _, scenario := range website.Scenarios {
		zapLogger.Debug("Scenario:", zap.String("ID", scenario.Id))
		scenarioFolderPath := filepath.Join(website.ScenariosFolderPath, scenario.Id)
		zapLogger.Debug("Scenario folder path:", zap.String("Path", scenarioFolderPath))
		err = resetVolatileFolder(scenarioFolderPath)
		if err != nil {
			zapLogger.Error(err.Error())
			return err
		}
		scenarioPage, scenarioPageMarshallErr := scenario.Marshal()
		if scenarioPageMarshallErr != nil {
			zapLogger.Error(scenarioPageMarshallErr.Error())
			return scenarioPageMarshallErr
		}
		scenarioPageWriteErr := ioutil.WriteFile(filepath.Join(scenarioFolderPath, "_index.md"), scenarioPage, os.ModePerm)
		if scenarioPageWriteErr != nil {
			zapLogger.Error(scenarioPageWriteErr.Error())
			return scenarioPageWriteErr
		}
		for _, notification := range scenario.Notifications {
			notificationFilePath := filepath.Join(scenarioFolderPath, fmt.Sprintf("%v.md", notification.Step.Position))
			notificationPage, notificationMarshallErr := notification.Marshal()
			if notificationMarshallErr != nil {
				zapLogger.Error(notificationMarshallErr.Error())
				return notificationMarshallErr
			}
			notificationPageWriteErr := ioutil.WriteFile(notificationFilePath, notificationPage, os.ModePerm)
			if notificationPageWriteErr != nil {
				zapLogger.Error(notificationPageWriteErr.Error())
				return notificationPageWriteErr
			}
		}
	}
	for _, notificationPattern := range website.NotificationsPatterns {
		notificationPatternFolderPath := filepath.Join(website.NotificationPatternsFolderPath, notificationPattern.Id)
		err = resetVolatileFolder(notificationPatternFolderPath)
		if err != nil {
			zapLogger.Error(err.Error())
			return err
		}
		notificationPatternPage, notificationPatternPageMarshallErr := notificationPattern.Marshal()
		if notificationPatternPageMarshallErr != nil {
			zapLogger.Error(notificationPatternPageMarshallErr.Error())
			return notificationPatternPageMarshallErr
		}
		notificationPatternPageWriteErr := ioutil.WriteFile(filepath.Join(notificationPatternFolderPath, "_index.md"), notificationPatternPage, os.ModePerm)
		if notificationPatternPageWriteErr != nil {
			zapLogger.Error(notificationPatternPageWriteErr.Error())
			return notificationPatternPageWriteErr
		}

	}
	return err
}

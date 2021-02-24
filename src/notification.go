package main

import (
	"gopkg.in/yaml.v3"
	"time"
)

type StepStruct struct {
	Scope    string `yaml:"scope"`
	Position int    `yaml:"position"`
	Sender   string `yaml:"sender"`
}

type NotificationConfig struct {
	ScenarioId  string `csv:"scenario_id"`
	Title       string `csv:"title"`
	Description string `csv:"description"`
	Updated     string `csv:"-"`
	Scope       string `csv:"scope"`
	Position    int    `csv:"position"`
	Sender      string `csv:"sender"`
	PatternId   string `csv:"pattern_id"`
}

type Notification struct {
	ScenarioId  string     `yaml:"-"`
	Title       string     `yaml:"title"`
	Description string     `yaml:"description"`
	Step        StepStruct `yaml:"step"`
	Updated     string     `yaml:"date"`
	PatternId   string     `yaml:"pattern_id"`
}

func (notification *Notification) Marshal() ([]byte, error) {
	notification.Updated = time.Now().Format("2006-01-02")
	webpageBytes, err := yaml.Marshal(notification)
	finalPage := append([]byte("---\n"), webpageBytes...)
	finalPage = append(finalPage, []byte("---\n\n")...)
	return finalPage, err
}

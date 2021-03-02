package main

import (
	"gopkg.in/yaml.v3"
	"time"
)

type PayloadExample struct {
	Number      int    `yaml:"number"`
	Example     string `yaml:"example"`
	Description string `yaml:"description"`
}

type WorkflowStep struct {
	ScenarioId          string `csv:"scenario_id" yaml:"-"`
	Title               string `csv:"title" yaml:"title"`
	Description         string `csv:"description" yaml:"description"`
	Updated             string `csv:"-" yaml:"date"`
	Scope               string `csv:"scope" yaml:"scope"`
	Position            int    `csv:"position" yaml:"position"`
	Sender              string `csv:"sender" yaml:"sender"`
	MandatoryProperties string `csv:"mandatory_properties" yaml:"mandatory_properties"`
	JsonPayload         string `csv:"json_payload" yaml:"json_payload"`
}

func (notification *WorkflowStep) Marshal() ([]byte, error) {
	notification.Updated = time.Now().Format("2006-01-02")
	webpageBytes, err := yaml.Marshal(notification)
	finalPage := append([]byte("---\n"), webpageBytes...)
	finalPage = append(finalPage, []byte("---\n\n")...)
	return finalPage, err
}

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

type NotificationConfig struct {
	ScenarioId          string `csv:"scenario_id"`
	Title               string `csv:"title"`
	Description         string `csv:"description"`
	Updated             string `csv:"-"`
	Scope               string `csv:"scope"`
	Position            int    `csv:"position"`
	Sender              string `csv:"sender"`
	MandatoryPayload    string `csv:"mandatory_payload"`
	PayloadExample1     string `csv:"payload_example_1"`
	PayloadDescription1 string `csv:"payload_description_1"`
	PayloadExample2     string `csv:"payload_example_2"`
	PayloadDescription2 string `csv:"payload_description_2"`
	PayloadExample3     string `csv:"payload_example_3"`
	PayloadDescription3 string `csv:"payload_description_3"`
	PayloadExample4     string `csv:"payload_example_4"`
	PayloadDescription4 string `csv:"payload_description_4"`
}

type WorkflowStep struct {
	ScenarioId       string           `yaml:"-"`
	Title            string           `yaml:"title"`
	Description      string           `yaml:"description"`
	Updated          string           `yaml:"date"`
	Scope            string           `yaml:"scope"`
	Position         int              `yaml:"position"`
	Sender           string           `yaml:"sender"`
	MandatoryPayload string           `yaml:"mandatory_payload"`
	PayloadExamples  []PayloadExample `yaml:"payload_examples"`
	//PayloadExample1     string `yaml:"payload_example_1"`
	//PayloadDescription1 string `yaml:"payload_description_1"`
	//PayloadExample2     string `yaml:"payload_example_2"`
	//PayloadDescription2 string `yaml:"payload_description_2"`
	//PayloadExample3     string `yaml:"payload_example_3"`
	//PayloadDescription3 string `yaml:"payload_description_3"`
	//PayloadExample4     string `yaml:"payload_example_4"`
	//PayloadDescription4 string `yaml:"payload_description_4"`
}

func (notification *WorkflowStep) Marshal() ([]byte, error) {
	notification.Updated = time.Now().Format("2006-01-02")
	webpageBytes, err := yaml.Marshal(notification)
	finalPage := append([]byte("---\n"), webpageBytes...)
	finalPage = append(finalPage, []byte("---\n\n")...)
	return finalPage, err
}

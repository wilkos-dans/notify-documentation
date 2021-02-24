package main

import (
	"gopkg.in/yaml.v3"
	"time"
)

type NotificationPattern struct {
	Id              string `csv:"id" yaml:"-"`
	Title           string `csv:"title" yaml:"title"`
	Description     string `csv:"description" yaml:"description"`
	PayloadExample1 string `csv:"payload_example_1" yaml:"payload_example_1"`
	PayloadExample2 string `csv:"payload_example_2" yaml:"payload_example_2"`
	PayloadExample3 string `csv:"payload_example_3" yaml:"payload_example_3"`
	PayloadExample4 string `csv:"payload_example_4" yaml:"payload_example_4"`
	Updated         string `csv:"-" yaml:"date"`
	Layout          string `yaml:"layout"`
}

func (notificationPattern *NotificationPattern) Marshal() ([]byte, error) {
	notificationPattern.Updated = time.Now().Format("2006-01-02")
	notificationPattern.Layout = "notification_pattern"
	webpageBytes, err := yaml.Marshal(notificationPattern)
	finalPage := append([]byte("---\n"), webpageBytes...)
	finalPage = append(finalPage, []byte("---\n\n")...)
	return finalPage, err
}

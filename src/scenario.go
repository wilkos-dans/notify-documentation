package main

import (
	"gopkg.in/yaml.v3"
	"time"
)

type Scenario struct {
	Id          string `yaml:"-" csv:"id"`
	Title       string `yaml:"title" csv:"title"`
	Description string `yaml:"description" csv:"description"`
	Actor1      string `yaml:"actor_1" csv:"actor_1"`
	Actor2      string `yaml:"actor_2" csv:"actor_2"`
	Layout      string `yaml:"layout"`
	Updated     string `yaml:"date"`
	Notifications []Notification `yaml:"-"`
}

func (scenario *Scenario) Marshal() ([]byte, error) {
	scenario.Updated = time.Now().Format("2006-01-02")
	scenario.Layout = "scenario"
	webpageBytes, err := yaml.Marshal(scenario)
	finalPage := append([]byte("---\n"), webpageBytes...)
	finalPage = append(finalPage, []byte("---\n\n")...)
	return finalPage, err
}

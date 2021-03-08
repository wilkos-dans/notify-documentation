package main

import (
	"gopkg.in/yaml.v3"
	"time"
)

type Pattern struct {
	Id                  string `csv:"pattern_id" yaml:"-"`
	Title               string `csv:"title" yaml:"title"`
	Description         string `csv:"description" yaml:"description"`
	Updated             string `csv:"-" yaml:"date"`
	MandatoryProperties string `csv:"mandatory_properties" yaml:"mandatory_properties"`
	OptionalProperties  string `csv:"optional_properties" yaml:"optional_properties"`
	//Layout              string `yaml:"layout"`
	//JsonPayload         string `csv:"json_payload" yaml:"json_payload"`
}

func (pattern *Pattern) Marshal() ([]byte, error) {
	pattern.Updated = time.Now().Format("2006-01-02")
	//pattern.Layout = "pattern"
	webpageBytes, err := yaml.Marshal(pattern)
	finalPage := append([]byte("---\n"), webpageBytes...)
	finalPage = append(finalPage, []byte("---\n\n")...)
	return finalPage, err
}

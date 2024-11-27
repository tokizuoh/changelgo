package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Target struct {
	Items []Item `yaml:"items"`
}

type Item struct {
	Link   string `yaml:"link"`
	Owner  string `yaml:"owner"`
	Repo   string `yaml:"repo"`
	File   string `yaml:"file"`
	Branch string `yaml:"branch"`
}

func read() (Target, error) {
	filename := "items.yml"
	data, err := os.ReadFile(filename)
	if err != nil {
		return Target{}, fmt.Errorf("Error reading to file: %v", err)
	}

	var target Target
	err = yaml.Unmarshal(data, &target)
	if err != nil {
		return Target{}, fmt.Errorf("Error decoding to yaml: %v", err)
	}

	return target, nil
}

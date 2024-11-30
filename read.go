package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func read() (Target, error) {
	filename := "items.yml"
	data, err := os.ReadFile(filename)
	if err != nil {
		return Target{}, fmt.Errorf("error reading to file: %v", err)
	}

	var target Target
	err = yaml.Unmarshal(data, &target)
	if err != nil {
		return Target{}, fmt.Errorf("error decoding to yaml: %v", err)
	}

	return target, nil
}

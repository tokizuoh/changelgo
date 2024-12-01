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

	var items Target
	err = yaml.Unmarshal(data, &items)
	if err != nil {
		return Target{}, fmt.Errorf("error decoding to yaml: %v", err)
	}

	return items, nil
}

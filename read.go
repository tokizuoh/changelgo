package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func read() (FileList, error) {
	filename := "items.yml"
	data, err := os.ReadFile(filename)
	if err != nil {
		return FileList{}, fmt.Errorf("error reading to file: %v", err)
	}

	var filelist FileList
	err = yaml.Unmarshal(data, &filelist)
	if err != nil {
		return FileList{}, fmt.Errorf("error decoding to yaml: %v", err)
	}

	return filelist, nil
}

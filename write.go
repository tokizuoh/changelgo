package main

import (
	"fmt"
	"os"
)

func write(rss string, filename string) error {
	dir := "./generated"
	path := dir + "/" + filename
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating to file: %v", err)
		}
	}

	err := os.WriteFile(path, []byte(rss), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

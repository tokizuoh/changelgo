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
			return fmt.Errorf("Error creating to file:", err)
		}
	}

	err := os.WriteFile(path, []byte(rss), 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file:", err)
	}

	return nil
}

package main

import (
	"fmt"
	"os"
)

func write(rss string) error {
	dir := "./generated"
	filename := dir + "/rss.xml"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Error creating to file:", err)
		}
	}

	err := os.WriteFile(filename, []byte(rss), 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file:", err)
	}

	return nil
}

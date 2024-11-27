package main

import (
	"fmt"
	"os"
)

func main() {
	response, err := fetch()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rss, err := generated(response)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dir := "./generated"
	filename := dir + "/rss.xml"
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating to file:", err)
			return
		}
	}

	err = os.WriteFile(filename, []byte(rss), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

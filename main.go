package main

import (
	"fmt"
	"os"
)

func main() {
	target, err := read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO: use target

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

	err = write(rss)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

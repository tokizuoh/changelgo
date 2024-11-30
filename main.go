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

	for _, item := range target.Items {
		response, err := fetch(item)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		rss, err := generate(item, response)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		filename := fmt.Sprintf("%s-%s-rss.xml", item.Owner, item.Repo)
		err = write(rss, filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

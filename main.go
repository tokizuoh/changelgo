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

	filelist := convert(target)

	for _, file := range filelist.Files {
		response, err := fetch(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		rss, err := generate(file, response)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		filename := fmt.Sprintf("%s-%s-rss.xml", file.Owner, file.Repo)
		err = write(rss, filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

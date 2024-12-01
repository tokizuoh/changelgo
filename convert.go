package main

import "strings"

func convert(target Target) FileList {
	var fileList FileList
	for _, item := range target.Items {
		parts := strings.Split(item.Link, "/")
		if len(parts) < 7 {
			continue
		}

		file := File{
			Link:   item.Link,
			Owner:  parts[3],
			Repo:   parts[4],
			Branch: parts[6],
			Name:   strings.Join(parts[7:], "/"),
		}
		fileList.Files = append(fileList.Files, file)
	}
	return fileList
}

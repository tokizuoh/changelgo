package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/gorilla/feeds"
)

func main() {
	response, err := fetch()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	feed := &feeds.Feed{
		Title: "swiftlang/swift CHANGELOG.md",
		Link: &feeds.Link{
			Href:   "https://github.com/swiftlang/swift/blob/main/CHANGELOG.md",
			Rel:    "alternate",
			Type:   "text/html",
			Length: "",
		},
		Description: "The commit history of the CHANGELOG.md in the swiftlang/swift repository.",
		Created:     time.Now(),
	}

	for _, edge := range response.Data.Repository.Object.History.Edges {
		message := edge.Node.Message
		limit := 75
		if utf8.RuneCountInString(message) > limit {
			runes := []rune(message)
			message = string(runes[:limit]) + "..."
		}
		re := regexp.MustCompile(`\s*(\n|\r)\s*`)
		cleanedTitle := re.ReplaceAllString(message, " ")

		const layout = "2006-01-02T15:04:05Z"
		pubDate, err := time.Parse(layout, edge.Node.CommittedDate)
		if err != nil {
			fmt.Println("Failed to parse date:", err)
			return
		}

		item := &feeds.Item{
			Title: cleanedTitle,
			Link: &feeds.Link{
				Href:   edge.Node.URL,
				Rel:    "alternate",
				Type:   "text/html",
				Length: "",
			},
			Description: edge.Node.Message,
			Created:     pubDate,
		}
		feed.Items = append(feed.Items, item)
	}

	rss, err := feed.ToRss()
	if err != nil {
		fmt.Printf("Failed to convert to rss: %v\n", err)
		return
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

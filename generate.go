package main

import (
	"fmt"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/gorilla/feeds"
)

func generate(file File, response Response) (string, error) {
	feed := &feeds.Feed{
		Title: fmt.Sprintf("%s/%s %s", file.Owner, file.Repo, file.Name),
		Link: &feeds.Link{
			Href:   file.Link,
			Rel:    "alternate",
			Type:   "text/html",
			Length: "",
		},
		// Description: "The commit history of the CHANGELOG.md in the swiftlang/swift repository.",
		Description: fmt.Sprintf("The commit history of the %s in the %s/%s repository.", file.Name, file.Owner, file.Repo),
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
			return "", fmt.Errorf("failed to parse date: %v", err)
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
		return "", fmt.Errorf("failed to convert to rss: %v", err)
	}

	return rss, nil
}

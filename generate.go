package main

import (
	"fmt"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/gorilla/feeds"
)

func generated(item Item, response Response) (string, error) {
	feed := &feeds.Feed{
		Title: fmt.Sprintf("%s/%s %s", item.Owner, item.Repo, item.File),
		Link: &feeds.Link{
			Href:   item.Link,
			Rel:    "alternate",
			Type:   "text/html",
			Length: "",
		},
		// Description: "The commit history of the CHANGELOG.md in the swiftlang/swift repository.",
		Description: fmt.Sprintf("The commit history of the %s in the %s/%s repository.", item.File, item.Owner, item.Repo),
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
			return "", fmt.Errorf("Failed to parse date:", err)
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
		return "", fmt.Errorf("Failed to convert to rss: %v\n", err)
	}

	return rss, nil
}

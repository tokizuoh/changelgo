package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/gorilla/feeds"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("GITHUB_TOKEN is not set")
		return
	}

	queryFile := "Query.graphql"
	queryBytes, err := os.ReadFile(queryFile)
	if err != nil {
		fmt.Printf("Failed to read query file: %v\n", err)
		return
	}
	query := string(queryBytes)

	// TODO: DI
	variables := map[string]interface{}{
		"owner":    "swiftlang",
		"name":     "swift",
		"first":    5,
		"filePath": "CHANGELOG.md",
	}

	body := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("Failed to marshal body: %v\n", err)
		return
	}

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(bodyBytes))
	if err != nil {
		fmt.Printf("Failed to create request: %v\n", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response: %v\n", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Non-OK HTTP status: %d\nResponse: %s\n", resp.StatusCode, string(respBody))
		return
	}

	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Printf("Failed to parse response: %v\n", err)
		return
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

	filename := "./generated/rss.xml"
	dir := "./generated"
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

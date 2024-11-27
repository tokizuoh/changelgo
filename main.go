package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
}

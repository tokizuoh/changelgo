package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetch(item Item) (Response, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return Response{}, fmt.Errorf("GITHUB_TOKEN is not set")
	}

	queryFile := "Query.graphql"
	queryBytes, err := os.ReadFile(queryFile)
	if err != nil {
		return Response{}, fmt.Errorf("Failed to read query file: %v\n", err)
	}
	query := string(queryBytes)

	// TODO: DI
	variables := map[string]interface{}{
		"owner":    item.Owner,
		"name":     item.Repo,
		"first":    5,
		"filePath": item.File,
	}

	body := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return Response{}, fmt.Errorf("Failed to marshal body: %v\n", err)
	}

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return Response{}, fmt.Errorf("Failed to create request: %v\n", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Response{}, fmt.Errorf("Failed to send request: %v\n", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, fmt.Errorf("Failed to read response: %v\n", err)
	}

	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("Non-OK HTTP status: %d\nResponse: %s\n", resp.StatusCode, string(respBody))
	}

	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return Response{}, fmt.Errorf("Failed to parse response: %v\n", err)
	}

	return response, nil
}

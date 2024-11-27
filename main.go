package main

import (
	"fmt"
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
}

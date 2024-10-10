package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func getPrettyJson(response *[]byte) (bytes.Buffer, error) {
	var prettyJson bytes.Buffer
	err := json.Indent(&prettyJson, *response, "", "  ") // Indent with 2 spaces
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("error formatting JSON: %w", err)
	}
	return prettyJson, nil
}

func viewResponse(response *[]byte) {
	// Unmarshal the JSON response
	var searchResp SearchResponse

	err := json.Unmarshal(*response, &searchResp)
	if err != nil {
		fmt.Println("error unmarshalling JSON response", err)
		return
	}

	for _, item := range searchResp.Items {
		fmt.Printf("Title: %s\nLink: %s\n\n", item.Title, item.Link)
	}
}

type SearchResponse struct {
	Items []SearchItem `json:"items"`
}

type SearchItem struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

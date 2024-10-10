package main

import (
	"fmt"
	"log"
	"os"
)

func printFile(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

func saveResponseToFile(response *[]byte) {
	const filename = "./resources/response.json"

	prettyJson, err := getPrettyJson(response)
	if err != nil {
		fmt.Println("Could not get pretty Json:", err)
		return
	}

	err = os.WriteFile(filename, prettyJson.Bytes(), 0644) // Write to file with read/write permissions
	if err != nil {
		println("error writing to file", err)
		return
	}
}

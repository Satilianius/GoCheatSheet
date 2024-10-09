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

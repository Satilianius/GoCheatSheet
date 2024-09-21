package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func printDefer() {
	printVariableDeferred()
	printDeferredStack()
	printFileContent("./resources/Hello.txt")
	deferInALoop()
}

func printVariableDeferred() {
	i := 0
	defer fmt.Printf("Value at the point of deferring: %d\n", i) // 0
	i++
	return
}

func printDeferredStack() {
	defer fmt.Println("Beginning of stack, executed last")
	defer fmt.Println("End of stack, executed first")
}

func deferInALoop() {
	nOfFiles := 3
	// To clean resources in the end of iteration - wrap defer calling code in a function
	for fileId := 0; fileId < nOfFiles; fileId++ {
		fileName := fmt.Sprintf("./resources/%d.txt", fileId)
		printFileContent(fileName)
	}
}

func printFileContent(filePath string) {

	// Open file
	pFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Close file on function end/exception
	defer func() {
		if err := pFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Do what you want with the file
	content, err := io.ReadAll(pFile)
	fmt.Printf("File content: %s\n", content)
}

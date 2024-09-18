package main

import (
	"fmt"
)

func main() {
	name := "Dima"
	fmt.Printf("Hello %s\n", name)

	fmt.Println("\n----------Times----------")
	printTimes()

	fmt.Println("\n----------Functions------")
	printFunctions()

	fmt.Println("\n----------Switch----------")
	printSwitch()
}

// TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

// Monzo recommends: https://go.dev/tour

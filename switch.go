package main

import (
	"fmt"
	"runtime"
)

func printSwitch() {
	printSimpleSwitch()
}

func printSimpleSwitch() {
	fmt.Print("Go runs on ")
	// Switch variable can be declared inside switch and be scoped by it
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

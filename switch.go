package main

import (
	"fmt"
	"runtime"
	"slices"
	"strconv"
)

func printSwitch() {
	printSimpleSwitch()
	fmt.Printf("Type switch: %s\n", returnTypeSwitch(42))
	printSwitchWithFunctions()
	// No switch expressions like return switch (v) { case "x" -> "result" }
}

func printSimpleSwitch() {
	fmt.Print("Go runs on ")
	// Switch variable can be declared inside switch and be scoped by it
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		// Breaks are implicit, if not needed - use `fallthrough`
	case "linux":
		fmt.Println("Linux.")
	// Multiple cases
	case "freebsd", "openbsd":
		fmt.Println("BSDs.")
	case "plan9": // Noop case, will do nothing
	default:
		fmt.Printf("%s.\n", os)
	}
}

func returnTypeSwitch(value interface{}) string {
	// Can declare another variable, which will have a define type inside the case
	switch castedValue := value.(type) {
	case int:
		return "int: " + strconv.Itoa(castedValue)
	case string:
		return "string: " + castedValue
	default:
		return "unknown"
	}
}

func printSwitchWithFunctions() {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"d", "e", "f"}
	value := "e"

	fmt.Print("Switch without variable, but with function: ")
	// switch without variable is equal to `switch true`
	switch {
	case slices.Contains(slice1, value):
		fmt.Println("slice1")
	case slices.Contains(slice2, value):
		fmt.Println("slice2")
	}
}

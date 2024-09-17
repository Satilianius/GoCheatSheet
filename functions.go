package main

import "fmt"

func printFunctions() {
	// Pass function by reference
	str1, str2 := stringsApply("one", "two", swap)
	fmt.Printf("Applying swap: str1: %s, str2: %s\n", str1, str2)

	// Pass functions as a lambda
	str1, str2 = stringsApply("one", "two", func(str1, str2 string) (string, string) { return str1 + "one", str2 + "two" })
	fmt.Printf("Applying lambda: str1: %s, str2: %s\n", str1, str2)

	// Invoking using interface
	fmt.Printf("Applying using interface: %s\n", applyConvertTo123(getStringValue))
}

// Declare function as parameter
func stringsApply(str1 string, str2 string, functionToApply func(string, string) (string, string)) (two string, one string) {
	return functionToApply(str1, str2)
}

// Can specify type for each or for all
func swap(str1, str2 string) (two string, one string) {
	// Can return tuples
	return str2, str1
}

// Passing functions by interface
// Declare a function interface: convert types take an int and return a string getStringValue.
type convert func(int) string

// getStringValue implements convert, returning x as string.
func getStringValue(x int) string {
	return fmt.Sprintf("%v", x)
}

// applyConvertTo123 passes 123 to convert func and returns quoted string.
func applyConvertTo123(fn convert) string {
	return fmt.Sprintf("%q", fn(123))
}

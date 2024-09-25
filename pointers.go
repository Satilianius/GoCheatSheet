package main

import "fmt"

func printPointers() {
	answer := 42
	pAnswer := &answer // & - get address of <variable>
	fmt.Printf("answer: %v\n", answer)
	fmt.Printf("pointer: %v\n", pAnswer)
	fmt.Printf("dereferenced pointer: %v\n", *pAnswer) // * - get value pointed by <pointer>

	cyberpunk := 2077
	var pCyberpunk *int // panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Printf("nil pointer: %v\n", pCyberpunk)
	pCyberpunk = &cyberpunk
	fmt.Printf("old cyberpunk: %v\n", cyberpunk)
	*pCyberpunk++ // modifying value through pointer
	fmt.Printf("new cyberpunk: %v\n", cyberpunk)

	// No pointer arithmetic
}

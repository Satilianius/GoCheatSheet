package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func printStructs() {
	printFieldAccess()
	printStructValuesCreation()
	printAnonymousStructs()
	printNestedStructs()
}

func printFieldAccess() {
	vertex := Vertex{1, 2} // Allocated on stack, deleted on function exit
	pVertex := &vertex

	fmt.Printf("vertex: %v, field X: %v\n", vertex, vertex.X)
	fmt.Printf("field X via struct pointer: %v\n", pVertex.X) // No `->`, no (*p).X!

	pCoordinate := &vertex.X // = &(vertex.X) `.` has precedence
	fmt.Printf("&vertex.X: %v\n", *pCoordinate)
	vertex.X = 42
	fmt.Printf("&(vertex.X): %v\n", *pCoordinate)
}

type Structure struct {
	Description string
}

func printStructValuesCreation() {
	structure1 := Structure{} // Fields initialised with default values
	structure2 := Structure{"All fields are given"}
	structure3 := Structure{Description: "Fields are given with names"}
	pStructure := &Structure{"Allocated on the heap"}

	pStructure = nil // Clean up

	fmt.Printf("structure1: %v, structure2: %v, structure3: %v, pStructure: %v\n",
		structure1, structure2, structure3, pStructure)
}

func printAnonymousStructs() {
	employee := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Bleh",
		lastName:  "Blah",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee: ", employee)
}

func printNestedStructs() {

}

// https://golangbot.com/structs/

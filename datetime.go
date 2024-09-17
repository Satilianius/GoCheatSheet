package main

import (
	"fmt"
	"time"
)

func printTimes() {
	// Zoned date times
	fmt.Printf("time.Now():       %s - Local time with timezone\n", time.Now().String())
	fmt.Printf("time.Now().UTC(): %s - UTC time\n", time.Now().UTC().String())
	fmt.Printf("time.Now(): %s - UTC time\n", time.Now().UTC().String())

	// Timestamps
	fmt.Printf("time.Now().Unix():       %d - Unix timestamp (secs form 01.01.1970 UTC)\n", time.Now().Unix())
	fmt.Printf("time.Now().UTC().Unix(): %d - Unix timestamp (secs form 01.01.1970 UTC)\n", time.Now().UTC().Unix())
}

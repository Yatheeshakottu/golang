package main

import (
	"fmt"
	"time"
)

func functionA(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "Function A completed"
}

func functionB(ch chan string) {
	time.Sleep(10 * time.Second)
	ch <- "Function B completed"
}

func main() {
	// Create channels to receive results from functions
	chA := make(chan string)
	chB := make(chan string)

	// Run functions concurrently using goroutines
	go functionA(chA)
	go functionB(chB)

	// Wait for both functions to complete or until a timeout (10 seconds in this case)
	select {
	case resultA := <-chA:
		fmt.Println(resultA)
	case resultB := <-chB:
		fmt.Println(resultB)
	case <-time.After(10 * time.Second):
		fmt.Println("Timeout: Both functions did not complete within 10 seconds.")
	}
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("First Goroutine")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Second Goroutine")
	}()

	wg.Wait()
	fmt.Println("All Goroutines completed")
}

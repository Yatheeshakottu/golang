package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func increment() {
	fmt.Println("Hello world ")
	defer wg.Done()
}
func decrement() {
	fmt.Println("welcome")
}
func main() {
	wg.Add(2)
	go increment()
	//wg.Wait()
	//wg.Add(1)
	go decrement()
	//	wg.Wait()

	wg.Wait()
}

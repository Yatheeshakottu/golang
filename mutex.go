package main

import (
	"fmt"
	"sync"
)

var count = 0
var wg sync.WaitGroup
var m sync.Mutex

func increment() {
	defer wg.Done()
	m.Lock()
	count++
	m.Unlock()
}
func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment()

		wg.Wait()
	}
	fmt.Println("count=", count)
}

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var Mutex = sync.RWMutex{}
var counter = 0

func hello() {
	fmt.Printf("Hello world=%v\n", counter)
	wg.Done()
	Mutex.RUnlock()
}
func increment() {
	counter++
	wg.Done()
	Mutex.Unlock()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		Mutex.RLock()
		go hello()
		Mutex.Lock()
		go increment()
	}
	wg.Wait()
}

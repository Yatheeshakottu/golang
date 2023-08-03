package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Display(str string) {

	for i := 0; i < 2; i++ {
		//	time.Sleep(1 * time.Second)
		wg.Add(1)
		fmt.Println(str)
		wg.Done()

	}
}

func main() {
	go Display("welcome")
	Display("Hello")
	wg.Wait()
}

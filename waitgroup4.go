package main

import (
	"fmt"
	"sync"
)

var x = 0

func incrementor(wg *sync.WaitGroup) {
	x = x + 1
	defer wg.Done()
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementor(&w)

		w.Wait()
	}
	fmt.Println("the final value of x is", x)
}

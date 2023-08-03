package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x += 1//x=x+1//x++
	fmt.Println(x)
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 100; i++ {
		w.Add(1)
		go increment(&w)
		w.Wait()
	}
	fmt.Println("The final value of x:", x)

}

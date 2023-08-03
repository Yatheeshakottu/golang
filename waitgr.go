// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func increment(wg *sync.WaitGroup) {
	fmt.Println("Hello world")
	defer wg.Done()
}
func decrement(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("welcome")

}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go increment(&wg)
	go decrement(&wg)

	wg.Wait()
}

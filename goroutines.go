package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
	defer wg.Done()
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
	defer wg.Wait()
}
func main() {

	b()
	wg.Add(1)
	go a()
	wg.Wait()
	/*go b()
	wg.Add(1)
	go a()
	Wg.Wait()*/    //output:aaaaaaaabbbbbbbbbbbbbbbaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa

}

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	values := []int{}
	values = append(values, 77)
	fmt.Println(values)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go abc(ch, &wg)
	fmt.Println("HELLO")
	fmt.Println(<-ch)
	go xyz("GOPLAY GROUND", ch, &wg)
	fmt.Println(<-ch)
	wg.Wait()
}
func abc(ch chan int, wg *sync.WaitGroup) {
	ch <- 64
	wg.Done()
}
func xyz(s string, ch chan int, wg *sync.WaitGroup) {
	fmt.Println(s)
	ch <- 652
	wg.Done()
}

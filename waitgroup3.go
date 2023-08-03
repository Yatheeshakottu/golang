package main

import (
	"fmt"
	"sync"
)

func send(wg *sync.WaitGroup, limitchannel chan int, stop int) {
	defer wg.Done()
	for i := 0; i <= stop; i++ {
		limitchannel <- i
	}
}
func receive(wg *sync.WaitGroup, limitchannel chan int, stop int) {
	defer wg.Done()
	for i := 0; i <= stop; i++ {
		fmt.Println(<-limitchannel)
	}
}
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	limitchannel := make(chan int)
	go send(wg, limitchannel, 4)
	go receive(wg, limitchannel, 4)
	wg.Wait()
}

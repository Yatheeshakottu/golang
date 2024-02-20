package main

import (
	"fmt"
	"sync"
	"time"
)

func channelA(ch chan string, wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	ch <- "channel A is completed"
	defer wg.Done()
}
func channelB(ch chan string, wg *sync.WaitGroup) {
	time.Sleep(10 * time.Second)
	ch <- "channel B is completed"
	defer wg.Done()
}
func main() {
	x := make(chan string)
	y := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go channelA(x, &wg)
	go channelB(y, &wg)
	go func() {
		wg.Wait()
		close(x)
		close(y)

	}()
	for {
		select {
		case a, ok := <-x:
			if ok {
				fmt.Print(a)
			}
		case b, ok := <-y:
			if ok {
				fmt.Print(b)
			}
		case <-time.After(10 * time.Second):
			fmt.Print("That two functions did not complete in 10 seconds")
		}
	}
}

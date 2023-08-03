package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(even, odd, quit)
	receive(even, odd, quit)
	fmt.Println("about to exit")

}
func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the evenchannel:", v)
		case v := <-o:
			fmt.Println("from the oddchannel:", v)
		case v := <-q:
			fmt.Println("from the quitchannel:", v)

		}
	}
}
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

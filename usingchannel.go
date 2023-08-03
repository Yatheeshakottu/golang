package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go foo(c)
	bar(c)
	fmt.Println("about to  exit")
}
func foo(c chan<- int) {
	c <- 21
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}

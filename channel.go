package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	c <- 40
	c <- 24
	fmt.Println(<-c)
	fmt.Println(<-c)
}

//Declaration of channels based on directions
//1. Bidirectional channel : chan T
//2. Send only channel: chan <- T
//3. Receive only channel: <- chan T

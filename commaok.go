package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 32
		close(c)
	}()
	i, v := <-c
	fmt.Println(i, v)
	i, v = <-c
	fmt.Println(i, v)
}

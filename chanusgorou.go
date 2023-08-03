package main

import (
	"fmt"
	"math/rand"
)

func Getvalues(values chan int) {
	emp := rand.Intn(21)
	values <- emp
}
func main() {
	x := make(chan int)
	go Getvalues(x)
	y := <-x
	fmt.Println("The values are:", y)

}

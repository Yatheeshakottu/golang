package main

import (
	"fmt"
)

func read() func(int) int {
	sum := 0
	return func(m int) int {
		sum += m
		return sum
	}
}
func main() {
	pos, neg := read(), read()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

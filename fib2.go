package main

import (
	"fmt"
)

func fibonacci(num int) {
	a := 1
	b := 2
	c := b
	fmt.Printf("series is:%d %d", a, b)
	for {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}
}
func main() {
	fibonacci(21)
	fibonacci(45)
	fibonacci(120)
}

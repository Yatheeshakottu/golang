package main

import (
	"fmt"
)

func gibonacci(num int) {
	a := 0
	b := 2
	c := b
	fmt.Printf("series is:%d %d", a, b)
	for {
		c = b
		b = a - b
		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}
}
func main() {
	gibonacci(24)
	gibonacci(77)
	gibonacci(220)

}

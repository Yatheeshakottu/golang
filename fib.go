package main

import (
	"fmt"
)

func sum(num int) {
	a := 1
	b := 2
	c := b
	fmt.Println("the series%d%d:", a, b)
	for {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
			break

		}

		a = c
		fmt.Printf("%d ", b)
	}

}
func main() {
	sum(23)
	sum(40)
}

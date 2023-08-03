package main

import (
	"fmt"
)

/*
	func swap(a, b int) (int, int) {
		a, b = b, a
		return a,b
	}

	func main() {
		a := 87
		b := 42
		fmt.Println(swap(a, b))
	}
*/
func swap() []int {
	a, b := 98, 53
	a, b = b, a
	return []int{a, b}
}
func main() {
	fmt.Println(swap())
}

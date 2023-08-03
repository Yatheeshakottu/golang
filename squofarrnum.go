package main

import (
	"fmt"
)

func add(x []int) int {
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i] * x[i]
	}
	return sum
}
func main() {
	//	a := []int{1, 2, 3, 4, 5}
	a := []int{9, 8, 2, 5, 7, 1, 6, 3, 4}
	fmt.Println(add(a))
}

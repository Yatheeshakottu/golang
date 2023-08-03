package main

import (
	"fmt"
)

func reverse(x []int) []int {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x); j++ {
			if x[i] < x[j] { //x[i]<x[j]{}
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	return x
}
func main() {
	x := []int{9, 876, 223, 1, 2, 76, 987653, 3468, 976, 0}

	fmt.Println(reverse(x))
}

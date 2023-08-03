package main

import (
	"fmt"
)

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func main() {
	fmt.Println(min(5, 17))
	fmt.Println(max(4, 28))
}

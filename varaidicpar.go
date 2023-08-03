package main

import (
	"fmt"
)

// zero or more values of a type int is called variadic parameter
func main() {
	x := sum(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(x)
}
func sum(x ...int) int {
	//fmt.Println(x)
	fmt.Printf("%d\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println(i, v)
	}
	//fmt.Println(sum)
	return sum
}

package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(x)
	fmt.Println("The length of the array is:", len(x))
	//fmt.Println(cap(x))
	fmt.Println(x[0])
	fmt.Println(x[5])
	fmt.Println(x[2])
	fmt.Println(x[4])
	fmt.Println(x[3])
	for x, y := range x {
		fmt.Println(x, y)
	}
}

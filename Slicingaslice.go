package main

import (
	"fmt"
)

func main() {
	x := []int{2, 3, 4, 56, 6, 7, 8}
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[:5])
	fmt.Println(x[:])
}

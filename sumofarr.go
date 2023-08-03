package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 49, 8}
	fmt.Println(x)
	sum := 0
	for i := 0; i < len(x); i++ {
		sum = sum + x[i]

	}
	fmt.Println(sum)
	avg := sum / len(x)
	fmt.Println("Average:", avg)
}

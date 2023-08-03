package main

import (
	"fmt"
)

func main() {
	var i, size int
	fmt.Print("Enter the array size:")
	fmt.Scan(&size)
	evenum := make([]int, size)
	fmt.Print("The  numbers in to the array are:")
	for i = 0; i < size; i++ {
		fmt.Scan(&evenum[i])
	}
	for i = 0; i < size; i++ {
		if evenum[i]%2 == 0 {

			fmt.Println(evenum[i])
		}

	}
	fmt.Println()
}

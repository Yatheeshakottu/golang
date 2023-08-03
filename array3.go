package main

import (
	"fmt"
)

func main() {
	var position, size int
	fmt.Print("Enter the size of the array:")
	fmt.Scanln(&size)
	smallarr := make([]int, size)
	fmt.Print("Enter the elements in to the array")
	for i := 0; i < size; i++ {
		fmt.Scanln(&smallarr[i])
	}
	smallest := smallarr[0]
	largest := smallarr[0]
	for i := 0; i < size; i++ {
		if largest > smallarr[i] {
			smallest = largest
			largest = smallarr[i]
		}
		if smallest > smallarr[i] && largest < smallest {
			smallest = smallarr[i]
			position = i
		}
	}
	fmt.Println("Second largest number in the list :", smallest)
	fmt.Println("The index position of the array  is:", position)
}

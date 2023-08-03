package main

import (
	"fmt"
)

func main() {
	var size, i, position int
	fmt.Print("Enter the size of array:")
	fmt.Scan(&size)
	smarr := make([]int, size)
	fmt.Print("Enter the elements in to array:")
	for i = 0; i < size; i++ {
		fmt.Scan(&smarr[i])
	}
	smallest := smarr[0]
	for i = 0; i < size; i++ {

		if smallest > smarr[i] {
			smallest = smarr[i]
			position = i
		}
	}
	fmt.Println("The smallest number in the array=", smallest)
	fmt.Println("The position of the number is:", position)
}

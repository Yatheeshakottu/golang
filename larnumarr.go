package main

import (
	"fmt"
)

func main() {
	var lagsize, position, i int
	fmt.Print("Enter the size of array:")
	fmt.Scan(&lagsize)
	lararr := make([]int, lagsize)
	fmt.Print("Enter the elements:")
	for i = 0; i < lagsize; i++ {
		fmt.Scan(&lararr[i])
	}
	largest := lararr[0]
	for i = 0; i < lagsize; i++ {
		if largest < lararr[i] {
			largest = lararr[i]
			position = i
		}
	}
	fmt.Println("The largest number:", largest)
	fmt.Println("The index position", position)
}

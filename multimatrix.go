package main

import (
	"fmt"
)

func main() {
	var i int
	var multiarr1 [5]int
	var multiarr2 [5]int
	var multiplicationarr [5]int
	fmt.Print("Enter the elements in to array:")
	for i = 0; i < 5; i++ {
		fmt.Scan(&multiarr1[i])
	}
	fmt.Print("Enter the elements into  array: ")
	for i = 0; i < 5; i++ {
		fmt.Scan(&multiarr2[i])
	}
	fmt.Println("The multiplication mateix is:")
	for i = 0; i < len(multiplicationarr); i++ {
		multiplicationarr[i] = multiarr1[i] * multiarr2[i]
		fmt.Println(multiplicationarr[i], "")
	}

	fmt.Println()
}

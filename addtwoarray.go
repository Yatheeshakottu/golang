package main

import (
	"fmt"
)

func main() {
	var addarray [7]int
	var addarray1 [7]int
	var addarray2 [7]int
	fmt.Print("Enter the elements in to array:")
	for i := 0; i < 7; i++ {
		fmt.Scan(&addarray1[i])
	}
	fmt.Print("Enter the elements into array2:")
	for i := 0; i < 7; i++ {
		fmt.Scan(&addarray2[i])
	}
	for i := 0; i < 7; i++ {

		addarray[i] = addarray1[i] + addarray2[i]

		fmt.Print(addarray[i], "\t")
	}
	fmt.Println()
}

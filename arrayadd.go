package main

import (
	"fmt"
)

func main() {
	//var arr1, arr2 int
	var arr1 [7]int
	var arr2 [7]int
	var arr3 [7]int
	fmt.Print("enter the first array elements")
	for i := 0; i < 7; i++ {
		fmt.Scan(&arr1[i])
	}
	fmt.Print("enter the second array elements")
	for i := 0; i < 7; i++ {
		fmt.Scan(&arr2[i])
	}
	for i := 0; i < 7; i++ {
		arr3[i] = arr1[i] + arr2[i]
		fmt.Print(arr3[i], "\t")
	}
	fmt.Println()
}

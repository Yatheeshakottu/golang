package main

import (
	"fmt"
)

func main() {
	var i, size int
	fmt.Print("Enter the size")
	fmt.Scan(&size)
	pcount := make([]int, size)
	fmt.Print("enter the elements to positive or negative number")
	for i = 0; i < size; i++ {
		fmt.Scan(&pcount[i])
	}
	poscount := 0
	negcount := 0
	for i = 0; i < size; i++ {
		if pcount[i] >= 0 {
			poscount++
		} else {
			negcount++
		}
	}
	fmt.Println("The positive count is:", poscount)
	fmt.Println("The negative count is:", negcount)
}

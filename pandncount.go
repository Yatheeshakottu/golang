package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, -8, -9, -3, -5, -3, 7, 8, 77}
	var pcount, ncount = 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 0 {
			pcount++
		} else {
			ncount++
		}
	}
	fmt.Println("The positive number count is", pcount)
	fmt.Println("The negative number count is:", ncount)

}

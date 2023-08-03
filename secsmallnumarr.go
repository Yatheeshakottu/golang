package main

import (
	"fmt"
)

func main() {
	var large1, large2 int
	arr := []int{2, 3, 4, 5, 6, 7}
	large1 = arr[0]
	for i := 0; i < len(arr); i++ {
		if large1 < arr[i] {
			large2 = large1
			large1 = arr[i]

		} else if large2 < arr[i] {
			//large2 = arr[i]
			large2 = large1
		}
	}
	fmt.Println(large2)
}

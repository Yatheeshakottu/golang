package main

import (
	"fmt"
)

func main() {
	var large1, large2, large3 int
	arr := []int{9, 42, 89, 1, 876, 15, 18, 398}
	for i := 0; i < len(arr); i++ {
		if large1 < arr[i] {
			large2 = large1
			large1 = arr[i]
		} else if large2 < arr[i] {
			large2 = arr[i]
		} else if large3 < arr[i] {
			large3 = arr[i]
		}
	}
	fmt.Println(large1, large2, large3)
}

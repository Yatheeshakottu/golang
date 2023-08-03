package main

import (
	"fmt"
)

func nonrepeatnum(arr []int) {
	for i := 0; i < len(arr); i++ {
		isnonrepeat := true
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] == arr[j] {
				isnonrepeat = false
				break
			}
		}
		if isnonrepeat {
			fmt.Print(arr[i], " ")
		}
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 4, 2, 1, 8, 9}
	nonrepeatnum(a)
}

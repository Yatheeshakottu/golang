package main

import (
	"fmt"
)

func main() {
	s := []int{5, 678, 987, 123}
	if len(s) == 0 {
		fmt.Println("slice is empty")
	} else {
		fmt.Println("slice is not empty")
	}

}

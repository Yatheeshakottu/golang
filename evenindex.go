package main

import (
	"fmt"
)

func main() {
	num := []int{10, 20, 40, 50, 30}
	fmt.Println(num)
	for i := range num {
		if i%2 == 0 {
			fmt.Println(num[i])
		}
	}
}

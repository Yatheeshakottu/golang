package main

import (
	"fmt"
)

func main() {
	y := [7]int{7, 17, 27, 37, 47, 57, 67}
	for i, v := range y {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", y)
}

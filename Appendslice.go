package main

import (
	"fmt"
)

func main() {
	x := []int{11, 22, 33, 44, 55, 66, 77, 88}
	fmt.Println(x)
	x = append(x, 99, 89, 90, 7, 67, 567)
	fmt.Println(x)
	y := append(x, 456, 678, 897, 675)
	x = append(x, y...)
	fmt.Println(x)
	x = append(x[:2], x[5:]...)
	fmt.Println(x)
}

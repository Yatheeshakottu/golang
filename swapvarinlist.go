package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	swap := reflect.Swapper(m)
	swap(0, 8)
	swap(1, 7)
	swap(2, 6)
	swap(3, 5)
	fmt.Println(m)
}

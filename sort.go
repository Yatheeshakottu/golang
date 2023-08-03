package main

import (
	"fmt"
	"sort"
)

func main() {
	y := []int{9, 76, 54, 6, 7, 8, 2, 1, 0, 5456, 67}
	s := []string{"yatheesha", "kottu"}
	fmt.Println(y)
	fmt.Println(s)
	sort.Ints(y)
	sort.Strings(s)
	fmt.Println("--------")
	fmt.Println(y)
	fmt.Println(s)

}

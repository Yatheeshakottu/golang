package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"yatheesha", "Prahakar", "sujatha", "teja", "swathi"}
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
	y := []int{9, 6, 5, 87, 65, 42, 1, 57}
	sort.Sort(sort.IntSlice(y))
	fmt.Println(y)
}

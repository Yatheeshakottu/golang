package main

import (
	"fmt"
)

func nonrepeatnum(s []int) {
	m := make(map[int]int)
	for _, val := range s {
		m[val]++
	}
	for _, val := range s {
		if m[val] == 1 {
			fmt.Print(val, " ")

		}
	}
}

func main() {
	a := []int{1, 5, 6, 7, 3, 1, 2, 2, 8, 5}
	nonrepeatnum(a)
}

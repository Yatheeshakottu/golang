package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	s[0], s[1] = s[1], s[0]
	s[2], s[3] = s[3], s[2]
	fmt.Println(s)
}
	
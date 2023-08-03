package main

import (
	"fmt"
)

func sort(s string) string {
	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		for j := 0; j < len(chars); j++ {
			if chars[i] < chars[j] {
				chars[i], chars[j] = chars[j], chars[i]
			}
		}
	}
	return string(chars)
}
func main() {
	s := "yatheesha"
	fmt.Println(sort(s))
}

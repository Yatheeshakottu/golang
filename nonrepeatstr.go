package main

import "fmt"

func nonRepeated(s string) string {
	m := make(map[rune]int)
	var res string
	for _, c := range s {
		m[c]++
	}
	for _, c := range s {
		if m[c] == 1 {
			res += string(c)
		}
	}
	return res
}

func main() {
	s := "hello world"
	fmt.Println(nonRepeated(s))
}

package main

import (
	"fmt"
)

func convertstring(input string) string {
	modified := make([]rune, 0, len(input)*2)
	for _, char := range input {
		if 'a' <= char && char <= 'z' {
			modified = append(modified, char, char-'a'+'A')
		} else if 'A' <= char && char <= 'z' {
			modified = append(modified, char, char-'A'+'a')

		} else {
			modified = append(modified, char)
		}
	}
	return string(modified)
}
func main() {
	input := "abc,ABC"
	modified := convertstring(input)
	fmt.Println(modified)
}

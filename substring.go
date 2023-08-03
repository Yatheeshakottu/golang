package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("enter the n value")
	fmt.Scan(&n)

	var s string
	fmt.Print("enter the s value")
	fmt.Scan(&s)

	count := 0
	for i := 0; i < n; i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			count += n - i // add all substrings that end with a vowel at index i
			for j := i + 1; j < n; j++ {
				if s[j] == 'a' || s[j] == 'e' || s[j] == 'i' || s[j] == 'o' || s[j] == 'u' {
					count++ // add all substrings that start with a vowel at index i and end with a vowel at index j
				}
			}
		}
	}

	fmt.Println(count)
}

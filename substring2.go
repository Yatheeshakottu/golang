package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	fmt.Scan(&s)

	count := 0
	vowels := "aeiouAEIOU"

	for i := 0; i < n; i++ {
		if strings.Contains(vowels, string(s[i])) {
			count += (n - i) * (i + 1)
		}
	}

	fmt.Println(count)
}

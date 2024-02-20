package main

import (
	"fmt"
	"strings"
)

func revstring(s string) string {
	words := strings.Fields(s)
	n := len(words)
	for i := 0; i < n/2; i++ {
		words[i], words[n-1-i] = words[n-1-i], words[i]
	}
	return strings.Join(words, " ")
}
func main() {
	input := "Hi hello good morning"
	fmt.Println(revstring(input))
}

/* 		words[n], words[n-1-i] = words[n-1-i], words[n].if i use like this instead of i ,n in a slice.
The issue arises from the fact that you're attempting to access an index n, which is out of bounds for the slice words.
 The valid indices for the slice words are from 0 to n-1.

In Go, the valid indices for a slice of length n are 0 through n-1.
 If you attempt to access an index outside this range, you will get a panic.*/

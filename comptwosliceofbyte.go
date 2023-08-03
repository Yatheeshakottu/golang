package main

import (
	"bytes"
	"fmt"
)

// by using compare
func main() {
	s1 := []byte{'Y', 'A', 'T', 'H', 'E', 'E', 'S', 'H', 'A'}
	s2 := []byte{'S', 'W', 'A', 'T', 'H', 'I'}
	res := bytes.Compare(s1, s2)
	if res == 0 {
		fmt.Println("slices are equal")

	} else {
		fmt.Println("slices are not equal")
	}
}

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	bs := []byte("◺")
	s := string(bs)
	fmt.Println(utf8.RuneCountInString(s)) // Output: 1
}

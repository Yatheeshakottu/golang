package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "ABC"
	b := "AB"
	c := strings.Contains(a, "AB")
	fmt.Println(b, c)
}

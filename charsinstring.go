package main

import (
	"fmt"
	"strings"
)

// by using contains
func main() {
	str1 := "Hello world"
	str2 := "Go playground"
	res1 := strings.Contains(str1, "hello")
	res2 := strings.Contains(str2, "Go")
	fmt.Println("Is  world present in string1", res1)
	fmt.Println("Is  go present in string1", res2)

}

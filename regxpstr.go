package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := regexp.MustCompile("el")
	fmt.Println(s.FindStringIndex("element"))
	fmt.Println(s.FindString("hello"))

}

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "welcome to homepage"
	x := "defer is delay the statement or function until the program is exit"
	z := "structure is a user defined data structure "
	res1 := strings.Replace(s, "m", "M", 5)
	res := strings.ReplaceAll(s, "to", "TO THE")
	res2 := strings.Replace(x, "e", "E", 12)
	res3 := strings.Replace(z, "t", "T", 10)
	fmt.Println(res1)
	fmt.Println(res)
	fmt.Println(res2)
	fmt.Println(res3)
}

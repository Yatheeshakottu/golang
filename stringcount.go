package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "hii hello everyonegood evening"
	y := "methods are  receiver type"
	z := "structure is a user defined datatype"
	res1 := strings.Count(x, "")
	res2 := strings.Count(y, "e")
	res3 := strings.Count(z, "t")
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}

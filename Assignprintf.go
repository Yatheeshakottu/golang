package main

import (
	"fmt"
)

func main() {
	type human int
	var x human
	var y int
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 77
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}

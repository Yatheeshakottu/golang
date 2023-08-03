package main

import (
	"fmt"
)

func main() {
	var n, x, y, z int
	if n == 0 {
		//fmt.Println("print n==0", x)
		return x
	}
	if n == 1 {
		//fmt.Println("print n==1", y)
		return y
	}
	z(gn) = y(gn-1) - x(gn-2)

	z = y - x
	x = y
	y = z
	fmt.Println(x, y, z)
}

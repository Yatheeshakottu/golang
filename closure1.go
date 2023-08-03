package main

import (
	"fmt"
)

func close() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	x := close()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	y := close()
	fmt.Println(y())
}

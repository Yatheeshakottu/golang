package main

import (
	"fmt"
)

func main() {
	a := 77
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
}

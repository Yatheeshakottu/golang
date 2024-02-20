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
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println("the address of b", 
	&b)
	fmt.Println(*&b)
	*b = 47
	fmt.Println(a)

}

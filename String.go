package main

import (
	"fmt"
)

func main() {
	s := "YATHEESHA"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	for i, v := range s {
		fmt.Printf("at the index position %d we have hex %#x\n", i, v)
	}
}

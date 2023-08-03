package main

import (
	"fmt"
)

const (
	a  = iota
	kb = 1 + (iota * 10)
	mb = 1 + (iota * 10)
	gb = 1 + (iota * 10)
)

func main() {
	fmt.Println(a)
	fmt.Printf("%d\t\t\t%b\n", a, a)

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)
}

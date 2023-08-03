package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "yathesha",
		last:  "kottu",
	}

	p2 := person{
		first: "kottu",
		last:  "yatheesha",
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p2.last)
	fmt.Println(p2.first, p2.last)

}

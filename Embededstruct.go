package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}
type secretagent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretagent{
		person: person{
			first: "yatheesha",
			last:  "kottu",
		},
		ltk: true,
	}

	p2 := person{
		first: "kottu",
		last:  "yatheesha",
	}
	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println(sa1.first, sa1.last, sa1.ltk)
	fmt.Println(p2.first, p2.last)

}

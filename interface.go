 package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}
type secreatagent struct {
	person
	ltk bool
}

func (s secreatagent) speak() {
	fmt.Println("I am", s.first, s.last)
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("i called human")
}

func main() {
	p1 := secreatagent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}
	p2 := secreatagent{
		person: person{
			first: "miss",
			last:  "moneypenny",
		},
		ltk: true,
	}
	p3 := person{
		first: "Dr.",
		last:  "james",
	}

	fmt.Println(p1)
	p1.speak()
	p2.speak()
	fmt.Println(p3)
}

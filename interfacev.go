package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}
type place struct {
	person
	living string
}

func (p person) speak() {
	fmt.Println("I am", p.name, "of", p.age)
}
func (p place) speak() {
	fmt.Println("I am ", p.name, "living in", p.living)

}

type human interface {
	speak()
}

func main() {
	p1 := person{
		name: "yatheesha",
		age:  21,
	}
	p2 := place{
		person: person{
			name: "swathi",
			age:  23,
		},
		living: "kottuvari palli",
	}
	p1.speak()
	p2.speak()
}

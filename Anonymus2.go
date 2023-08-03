package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{

		first: "james",
		last:  "bond",
		age:   34,
	}
	fmt.Println(p1)
}

package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"james":            32,
		"Miss Monneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["james"])
	if _, ok := m["miky"]; ok {
		fmt.Println(m["miky"])

	}
}

package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"yatheesha": 21,
		"swathi":    23,
	}
	fmt.Println(m)
	fmt.Println(m["swathi"])
	m["teja"] = 26
	for k, v := range m {
		fmt.Println(k, v)
	}
	xi := []int{1, 2, 3, 4, 5, 6}
	for i, v := range xi {
		fmt.Println(i, v)

	}
}

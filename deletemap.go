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
	delete(m, "swathi")
	fmt.Println(m)
	fmt.Println(m["teja"])
	fmt.Println(m["swathi"])
	if v, ok := m["yatheesha"]; ok {
		fmt.Println(v)
		delete(m, "yatheesha")
	}
	fmt.Println(m)

}

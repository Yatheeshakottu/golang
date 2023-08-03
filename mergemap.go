package main

import (
	"fmt"
)

func main() {
	var employee = map[string]int{"yatheesha": 10, "swathi": 12}
	employee = make(map[string]int)
	//	var student = map[string]int{"teja": 23, "sujatha": 43}
	//for i, k := range student {
	//employee[i] = k
	//}
	fmt.Println(employee)
}

package main

import (
	"fmt"
)

func main() {
	x := make(map[string]int)
	x["A"] = 7
	x["B"] = 6
	x["C"] = 5
	x["D"] = 4
	x["E"] = 3
	x["F"] = 2
	x["G"] = 1
	x["H"] = 0
	y := make(map[string]int)
	for i, v := range x {
		y[i] = v
	}
	for i, v := range y {
		fmt.Printf("Index: %s value: %d \n", i, v)
	}
}

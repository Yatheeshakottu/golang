package main

import (
	"fmt"
)

func main() {
	var i, j int
	fmt.Println("Multiplication table")
	for i = 2; i <= 20; i++ {
		for j = 1; j <= 10; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
		fmt.Println("-------------------")

	}
}

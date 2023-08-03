package main

import (
	"fmt"
)

func main() {
	var i, j int

	fmt.Print("Enter the elements in to the i")
	fmt.Scan(&i)
	for j = 1; j < i; j++ {
		j = i - 1

	}
	fmt.Println(j)
}

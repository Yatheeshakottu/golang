package main

import (
	"fmt"
)

func main() {
	var factorsnum, i int
	fmt.Print("Enter the number to find factors:")
	fmt.Scan(&factorsnum)
	for i = 1; i <= factorsnum; i++ {
		if factorsnum%i == 0 {
			fmt.Println(i)
		}
	}
}

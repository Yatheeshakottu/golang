package main

import (
	"fmt"
)

func main() {
	var i, j, rows, number int
	fmt.Print("enter the rows")
	fmt.Scan(&rows)
	for i = 1; i <= rows; i-- {
		for j = 1; j < i; j-- {

			fmt.Printf("%d\t", number)
			number++
		}
		fmt.Println()
	}
}

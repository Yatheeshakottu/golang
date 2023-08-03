package main

import (
	"fmt"
)

func main() {
	var rows int
	var number int
	fmt.Print("enter the rows")
	fmt.Scanln(&rows)
	fmt.Print("enter the starting number")
	fmt.Scanln(&number)
	for i := 0; i <= rows; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d", number)
			number++
		}

		fmt.Println()
	}
}

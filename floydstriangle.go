package main

import (
	"fmt"
)

func main() {
	var i, j, rows int
	var number int
	fmt.Print("enter total rows to pritnt floyds triangle=")
	fmt.Scanln(&rows)
	fmt.Print("enter the starting number=")
	fmt.Scanln(&number)

	fmt.Println("floyd's triangle")
	for i = 1; i <= rows; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d", number)
			number++
		}
		fmt.Println()
	}
}

package main

import "fmt"

func main() {
	var i, j, rows, number int
	fmt.Print("Enter the rows")
	fmt.Scan(&rows)
	fmt.Print("Enter the starting num")
	fmt.Scan(&number)

	for i = 0; i <= rows; i++ {
		for j = 0; j <= i; j++ {
			fmt.Printf(" %d ", number)

			number++
		}
		fmt.Println()
	}
}

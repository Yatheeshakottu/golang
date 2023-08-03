package main

import (
	"fmt"
)

func main() {
	var firstdigit, number int
	fmt.Print("Enter the digits to find the first digit")
	fmt.Scan(&number)
	firstdigit = number
	for firstdigit >= 10 {
		firstdigit = firstdigit / 10

	}
	fmt.Println(firstdigit)
}

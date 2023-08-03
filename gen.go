package main

import (
	"fmt"
)

func main() {
	var genNumber, sum, remainder int

	fmt.Print("Enter the Number to find Generic Root = ")
	fmt.Scanln(&genNumber)

	for genNumber >= 10 {
		for sum = 0; genNumber > 0; genNumber = genNumber / 10 {
			remainder = genNumber % 10
			sum = sum + remainder
		}
		genNumber = sum
	}

	fmt.Println("The Generic Root of this Number = ", genNumber)

}

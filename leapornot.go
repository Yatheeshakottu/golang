package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Print("Enter the year to find leap year or not:")
	fmt.Scan(&year)
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("It is a leap year")
	} else {
		fmt.Println("It is not a leap year")
	}
}

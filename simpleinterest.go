package main

import (
	"fmt"
)

// simple interest=(principal amount*rate of interest*years)/100
func main() {
	var simplei, pamount, rateinterest, years float64
	fmt.Print("Enter the principal amount:")
	fmt.Scan(&pamount)
	fmt.Print("Enter the interest rate:")
	fmt.Scan(&rateinterest)
	fmt.Print("Enter the  years:")
	fmt.Scan(&years)
	simplei = (pamount * rateinterest * years) / 100
	fmt.Println(simplei)
}

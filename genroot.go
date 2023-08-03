package main

import (
	"fmt"
)

func main() {
	var gennum, sum, remainder int
	fmt.Print("Enter the number to find generic number")
	fmt.Scan(&gennum)
	for gennum >= 10 {
		for sum = 0; gennum > 0; gennum = gennum / 10 {
			remainder = gennum % 10
			sum = sum + remainder
		}
		if sum >= 10 {
			gennum = sum
		} else {
			fmt.Println(sum)

		}
	}
}

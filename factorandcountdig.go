package main

import (
	"fmt"
)

func main() {
	var factornum int
	fmt.Print("Enter the value to find a factors:")
	fmt.Scanln(&factornum)
	for i := 1; i <= factornum; i++ {
		if factornum%i == 0 {
			fmt.Printf(" %d", i)
		}
	}
	count := 0
	for count = 0; factornum > 0; factornum = factornum / 10 {
		count++
	}
	fmt.Println(" \n The digits in anumber is :", count)
}

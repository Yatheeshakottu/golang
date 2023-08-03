package main

import (
	"fmt"
)

func main() {
	var primenum, primecount int
	primecount = 0
	fmt.Print("Enter the number to find prime or not:")
	fmt.Scanln(&primenum)
	for i := 2; i < primenum; i++ {
		if primenum%i == 0 {
			primecount++
			break
		}
	}
	if primenum != 1 && primecount == 0 {
		fmt.Println("It is aprime number=", primenum)
	} else {
		fmt.Println("It is not a prime number=", primenum)
	}

}

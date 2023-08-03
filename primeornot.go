package main

import (
	"fmt"
)

func main() {
	var primcount, primnum int
	primcount = 0
	fmt.Print("Enter the  number to find the prime number=")
	fmt.Scanln(&primnum)

	for i := 2; i < primnum; i++ {
		if primnum%i == 0 {
			primcount++
			break
		}
	}
	if primcount == 0 && primnum != 1 {
		fmt.Println(primnum, "it is prime number")
	} else {
		fmt.Println(primnum, "it is not a prime number")
	}
}

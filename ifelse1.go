package main

import (
	"fmt"
)

func main() {
	if x := 9; x < 0 {
		fmt.Println("it is negative number", x)
	} else if x < 10 {
		fmt.Println("9 is less than 10")
	} else {
		fmt.Println("it is a multiple of 3")
	}
	if 7%3 == 0 {
		fmt.Println("it is divisible")
	} else {
		fmt.Println("it is not divisible by  3 ")
	}
	if 10%2 == 0 {
		fmt.Println("it is divisible by 2")
	}
}

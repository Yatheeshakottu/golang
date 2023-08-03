package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Print("Enter the values to find divisible by:")
	fmt.Scan(&i)
	if i%5 == 0 && i%7 == 0 {
		fmt.Println("It is divisible by 5 and 7")

	} else {
		fmt.Println("It is not divisible by 5 and 7")
	}
}

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter the number to find positive or negative:")
	fmt.Scan(&num)
	if num >= 0 {
		fmt.Println("It is a positive number:", num)
	} else {
		fmt.Println("it is a negative number:", num)
	}

}

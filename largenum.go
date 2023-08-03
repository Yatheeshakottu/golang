package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Enter the first number")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number")
	fmt.Scan(&num2)

	if num1 > num2 {

		fmt.Println("It is not correct")
	} else {
		fmt.Println("It is correct")
	}
}

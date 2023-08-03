package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	hcf := findHCF(num1, num2)

	fmt.Printf("The highest common factor of %d and %d is %d", num1, num2, hcf)
}

func findHCF(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

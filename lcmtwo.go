package main

import "fmt"

// Function to find the GCD of two numbers
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Function to find the LCM of two numbers
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)//6,.mn Calculate the LCM using the formula LCM(a, b) = (a * b) / GCD(a, b)
   
}

func main() {
	var a, b int
	fmt.Print("Enter two integers: ")
	fmt.Scanln(&a, &b)
	fmt.Printf("LCM of %d and %d is %d\n", a, b, lcm(a, b))
}

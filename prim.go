package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Println(n, "is a prime number.")
	} else {
		fmt.Println(n, "is not a prime number.")
	}
}

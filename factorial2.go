package main

import (
	"fmt"
)

func main() {
	var factorial, i, factorialnum int
	factorial = 1
	fmt.Print("Enter the number to find factorial:")
	fmt.Scanln(&factorialnum)
	for i = 1; i <= factorialnum; i++ {
		factorial = factorial * i
	}
	fmt.Println(factorial)
}

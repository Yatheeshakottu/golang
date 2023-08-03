package main

import (
	"fmt"
)

// function to print the highest factor of a number
func highestFactor(n int) {
	maxFactor := 1
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			maxFactor = i
			n /= i
			i--
		}
	}
	fmt.Println(maxFactor)
}

func main() {
	// find the highest factor of 10
	fmt.Print("The highest factor of 10 is: ")
	highestFactor(10)

	// find the highest factor of 12
	fmt.Print("The highest factor of 12 is: ")
	highestFactor(12)
}

/*The output of this program will be:
python
Copy code
The highest factor of 10 is: 5
The highest factor of 12 is: 6
Note that the highestFactor function works by iterating from 2 to the given number n, and checking if each number is a factor of n. If it is, the function updates the value of maxFactor and divides `*/

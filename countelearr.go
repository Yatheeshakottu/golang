package main

import (
	"fmt"
)

func main() {
	m := 123456789
	count := 0
	for m > 0 {
		m = m / 10
		count++

	}
	fmt.Println("The count of the number is ", count)
}

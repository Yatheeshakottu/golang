package main

import (
	"fmt"
)

func main() {
	var largest, smallest int
	var maxposition, minposition int

	x := []int{12, 34, 1, 87, 56, 346, 87654, 4321}
	largest = x[0]
	smallest = x[0]
	for i := 0; i < len(x); i++ {

		if largest < x[i] {
			largest = x[i]
			maxposition = i
		}
		if smallest > x[i] {
			smallest = x[i]
			minposition = i
		}
	}
	fmt.Println("\nThe Largest Number in this lgsmArr   = ", largest)
	fmt.Println("The Index Position of Largest Number = ", maxposition)
	fmt.Println("\nThe Smallest Number in this lgsmArr   = ", smallest)
	fmt.Println("The Index Position of Smallest Number = ", minposition)

}

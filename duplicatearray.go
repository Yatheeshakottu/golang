package main

import (
	"fmt"
)

func main() {
	var dupsize, i int
	fmt.Print("Enter the array size:")
	fmt.Scan(&dupsize)
	duparr := make([]int, dupsize)
	fmt.Print("Enter the elements in to the array:")

	for i = 0; i < dupsize; i++ {
		fmt.Scan(&duparr[i])

	}
	dupcount := 0
	for i = 0; i < dupsize; i++ {
		for j := i + 1; j < dupsize; j++ {
			if duparr[i] == duparr[j] {
				dupcount++
				break
			}
		}
	}
	fmt.Println("The total number of duplicate elements:", dupcount)
}

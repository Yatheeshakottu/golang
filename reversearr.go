package main

import (
	"fmt"
)

func main() {
	var size, i, j int
	fmt.Print("Enter the array size:")
	fmt.Scan(&size)
	actnum := make([]int, size)
	revarr := make([]int, size)

	fmt.Print("Enter the numbers in to array: ")
	for i = 0; i < size; i++ {
		fmt.Scan(&actnum[i])

	}
	j = 0
	for i = size - 1; i >= 0; i-- {
		revarr[j] = actnum[i]
		j++
	}

	fmt.Println("The result of areverse array=", revarr)
}

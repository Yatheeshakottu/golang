package main

import "fmt"

func main() {
	// Create a slice with some repeated elements
	slice := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1, 6, 7}

	// Create a map to store the unique elements
	uniqueElements := make(map[int]bool)

	// Iterate over the slice and add each element to the map
	// If the element is already in the map, it will be ignored
	for _, element := range slice {
		uniqueElements[element] = true
	}

	// Print the unique elements
	//fmt.Println("Unique elements:", uniqueElements)
	fmt.Println(uniqueElements)
}

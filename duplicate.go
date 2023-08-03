package main

import "fmt"

func main() {
	var dupsize, i int

	fmt.Print("Enter the Even Array Size = ")
	fmt.Scan(&dupsize)

	dupArr := make([]int, dupsize)

	fmt.Print("Enter the Even Array Items  = ")
	for i = 0; i < dupsize; i++ {
		fmt.Scan(&dupArr[i])
	}
	dupcount := 0
	for i = 0; i < dupsize; i++ {
		for j := i + 1; j < dupsize; j++ {
			if dupArr[i] == dupArr[j] {
				dupcount++
				break
			}
		}
	}
	fmt.Println("\nThe Total Number of Duplicates in dupArr = ", dupcount)
}

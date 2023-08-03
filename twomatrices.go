package main

import (
	"fmt"
)

func main() {
	var i, j, rows, columns int
	var addmat1 [10][10]int
	var addmat2 [10][10]int
	var additionmat [10][10]int

	fmt.Print("enter the rows and columns=\n")
	fmt.Scan(&rows, &columns)

	fmt.Print("Enter the first matrix=\n")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&addmat1[i][j])
		}
	}
	fmt.Print("ENTER THE SECOND MATRICES =\n")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&addmat2[i][j])
		}
	}
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			additionmat[i][j] = addmat1[i][j] + addmat2[i][j]

		}
	}
	fmt.Println("The sum of two matrices =")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Print(additionmat[i][j], "\t")

		}
		fmt.Println()
	}

}

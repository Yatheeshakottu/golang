package main

import "fmt"

func main() {
	var i, j, rows, columns int
	var addmat [10][10]int
	var addmat2 [10][10]int
	var addmat3 [10][10]int
	fmt.Print("enter the rows and columns")
	fmt.Scan(&rows, &columns)
	fmt.Print("enter the first matrix")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&addmat[i][j])
		}
	}
	fmt.Print("enter the second matrix")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&addmat2[i][j])
		}
	}

	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			addmat3[i][j] = addmat[i][j] + addmat2[i][j]
		}
	}
	fmt.Print("Addition of the two matrix are\n")
	for i = 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Print(addmat3[i][j], "\t")
		}

		fmt.Println()
	}
}

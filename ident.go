package main

import "fmt"

func main() {
	var identmat [10][10]int
	var rows, columns int
	fmt.Print("Enter the rows and columns")
	fmt.Scan(&rows, &columns)
	fmt.Print("Enter elements to matrix")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Scan(&identmat[i][j])

		}
	}
	y := 1
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if identmat[i][j] != 1 && identmat[j][i] != 0 {
				y = 0
				break
			}
		}
	}
	if y == 1 {
		fmt.Println("It is a identity matrix")
	} else {
		fmt.Println("It is not a identity matix")
	}
}

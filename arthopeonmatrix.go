package main

import (
	"fmt"
)

func main() {
	var rows, columns int
	var gomat1 [10][10]int
	var gomat2 [10][10]int
	fmt.Print("Enter the rows and columns:")
	fmt.Scan(&rows, &columns)
	fmt.Print("Enter the first matrix elements:")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Scan(&gomat1[i][j])
		}
	}
	fmt.Print("Enter the second matrix elements:")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Scan(&gomat2[i][j])
		}
	}
	fmt.Print("ADD\tsub\tMUl\tdiv\tmod\t")
		for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Print("\n", gomat1[i][j]+gomat2[i][j], "\t")
			fmt.Print(gomat1[i][j]-gomat2[i][j], "\t")
			fmt.Print(gomat1[i][j]*gomat2[i][j], "\t")

			fmt.Print(gomat1[i][j]/gomat2[i][j], "\t")

			fmt.Print(gomat1[i][j]%gomat2[i][j], "\t")

		}
	}
}

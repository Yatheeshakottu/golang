package main

import (
	"fmt"
)

func main() {
	var rows, columns int
	var mat1 [10][10]int
	var mat2 [10][10]int
	fmt.Print("enter the rows and columns:")
	fmt.Scanln(&rows, &columns)
	fmt.Print("enter the first matrix:")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Scan(&mat1[i][j])
		}
	}
	fmt.Print("Enter the second matrix:")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Scan(&mat2[i][j])
		}
	}
	fmt.Print("ADD\tSUB\tMUL\tDIV\tMOd")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Print("\n", mat1[i][j]+mat2[i][j], "\t")
			fmt.Print(mat1[i][j]-mat2[i][j], "\t")
			fmt.Print(mat1[i][j]*mat2[i][j], "\t")
			fmt.Print(mat1[i][j]/mat2[i][j], "\t")
			fmt.Print(mat1[i][j]%mat2[i][j], "\t")
		}
	}
}

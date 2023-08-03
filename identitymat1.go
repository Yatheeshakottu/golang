package main

import "fmt"

func main() {
	var rows, columns, i, j int

	var identMat [10][10]int

	fmt.Print("Enter the Matrix Rows and Columns = ")
	fmt.Scan(&rows, &columns)

	fmt.Print("Enter the Matrix Items = ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&identMat[i][j])
		}
	}
	flag := 1
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			if (i == j) && (identMat[i][j] != 1) {
				flag = 0
			} else if (i != j) && (identMat[i][j] != 0) {
				flag = 0
			}
		}
	}
	if flag == 1 {
		fmt.Println("It is an Idenetity Matrix")
	} else {
		fmt.Println("It is Not an Idenetity Matrix")
	}
}

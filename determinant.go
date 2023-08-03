package main

import (
	"fmt"
)

func main() {
	var i, j int
	var determat [2][2]int
	fmt.Print("Enter the elements to the matrix:")
	for i = 0; i < 2; i++ {
		for j = 0; j < 2; j++ {

			fmt.Scan(&determat[i][j])
		}
	}
	determinant := (determat[0][0]*determat[1][1] - determat[0][1]*determat[1][0])
	fmt.Println("The determinant of the 2*2 matrix is =", determinant)
}

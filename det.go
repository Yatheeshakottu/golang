package main

import (
	"fmt"
)

func main() {
	var determat [2][2]int
	fmt.Print("enter the numbers:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&determat[i][j])

		}
	}
	determinant := (determat[0][0]*determat[1][1] - determat[0][1]*determat[1][0])
	fmt.Println(determinant)
}

package main

import (
	"fmt"
)

func main() {
	var detmat [2][2]int
	fmt.Print("Enter the values")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&detmat[i][j])
		}
	}
	determinant := (detmat[0][0]*detmat[1][1] - detmat[0][1]*detmat[1][0])
	fmt.Println(determinant)
}

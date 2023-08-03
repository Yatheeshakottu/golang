package main

import (
	"fmt"
)

func main() {
	var i, j int
	var idenmat [10][10]int
	matrix1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	s := 1
	for i = 0; i < len(matrix1); i++ {
		for j = 0; j < len(matrix2); j++ {
			if (i == j) && (idenmat[i][j] != 1) {
				s = 0
			} else if (i != j) && (idenmat[i][j] != 0) {
				s = 0
			}
		}
	}
	if s == 1 {
		fmt.Println("It is a identity matrix")
	} else {
		fmt.Println("IT is not a identity matrix")
	}

}

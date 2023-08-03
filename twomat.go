package main

import (
	"fmt"
)

func main() {
	mat1 := [3][4]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 4, 5}}
	mat2 := [3][4]int{{8, 1, 5, 2}, {6, 2, 3, 4}, {1, 2, 3, 5}}
	var mat3 [3][4]int
	var mat4 [3][4]int
	var mat5 [3][4]int
	var mat6 [3][4]int
	var mat7 [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			mat3[i][j] = mat1[i][j] + mat2[i][j]
			mat4[i][j] = mat1[i][j] - mat2[i][j]
			mat5[i][j] = mat1[i][j] * mat2[i][j]
			mat6[i][j] = mat1[i][j] % mat2[i][j]
			mat7[i][j] = mat1[i][j] / mat2[i][j]
		}
	}
	fmt.Println(mat1)
	fmt.Println(mat2)
	fmt.Println(mat3)
	fmt.Println(mat4)
	fmt.Println(mat5)
	fmt.Println(mat6)
	fmt.Println(mat7)
}

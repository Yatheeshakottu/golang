package main

import "fmt"

func main() {
	var i, j int
	var det [2][2]int
	fmt.Print("enter the vAlues")
	for i = 0; i < 2; i++ {
		for j = 0; j < 2; j++ {
			fmt.Scan(&det[i][j])
		}
	}
	determinant := det[0][0]*det[1][1] - det[0][1]*det[1][0]
	fmt.Println(determinant)

}

package main

import (
	"fmt"
)

func main() {
	var x [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			x[i][j] = i + j
		}
	}
	fmt.Println(x)
}

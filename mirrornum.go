package main

import (
	"fmt"
)

func main() {
	//s := "yatheesha"
	var rows int = 10
	for i := rows; i > 0; i-- {
		for j := rows - i; j > 0; j-- {
			fmt.Printf(" ")
		}
		for k := 1; k < i; k++ {
			fmt.Printf("%d", k)
		}
		fmt.Println()
	}
}

package main

import (
	"fmt"
)

func main() {
	// rows=5
	for i := 5; i >= 1; i-- {
		for j := 1; j <= 5-i; j++ {
			fmt.Print(" ")
		}
		for k := i; k <= 2*i-1; k++ {
			fmt.Print("7")
		}
		for k := 0; k < i-1; k++ {
			fmt.Print("7")
		}
		fmt.Println("")

	}
}

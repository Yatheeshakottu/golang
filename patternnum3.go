package main

import (
	"fmt"
)

func main() {
	var x, n int = 0, 5
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ")
			x = x + n - i
			fmt.Print(x)
		}
		fmt.Println()
	}
}

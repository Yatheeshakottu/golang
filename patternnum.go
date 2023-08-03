package main

import (
	"fmt"
)

func main() {
	var x, n int = 0, 5
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			x++
			if x == 10 {
				x = 1
			}
			fmt.Print(x)
		}
		fmt.Println()
	}
}

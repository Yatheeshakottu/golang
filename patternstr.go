package main

import (
	"fmt"
)

func main() {
	s := "yatheesha"
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(" ")
			fmt.Printf("%c", s[i])
		}

		fmt.Println()
	}
}


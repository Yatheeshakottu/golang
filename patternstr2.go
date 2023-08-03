package main

import (
	"fmt"
)

func main() {
	s := "yatheesha"

	for i := 0; i <= len(s); i++ {
		for j := 0; j < len(s)-i; j++ {
			fmt.Print(" ")
			fmt.Printf("%c", s[j])
		}
		fmt.Println()
	}
}

/*package main

import "fmt"

func main() {
	s := "yatheesha"
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i; j++ {
			fmt.Printf("%c ", s[i])
		}
		fmt.Println()
	}
}
*/

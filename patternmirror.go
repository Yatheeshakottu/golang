package main

import (
	"fmt"
)

func main() {

	for i := 10; i > 0; i-- {
		for j := 10 - i; j > 0; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Printf("%d", i)
			//fmt.Printf("%d",k)
		}
		fmt.Println()
	}
}

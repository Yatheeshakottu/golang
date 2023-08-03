package main

import (
	"fmt"
)

func main() {
	x := []int{2, 3, 3, 4, 5, 6, 7, 7, 7, 4, 3, 5}
	ecount := 0
	ocount := 0
	for i := 0; i < len(x); i++ {
		if i%2 == 0 {
			ecount++

		} else {
			ocount++

		}
	}
	fmt.Println("The even count:", ecount)
	fmt.Println("The odd count:", ocount)

}

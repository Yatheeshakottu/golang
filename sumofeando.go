package main

import (
	"fmt"
)

func main() {
	//var enum, i int
	enum := []int{1, 3, 2, 4, 5, 6, 7, 8, 90, 101}
	etotal := 0
	ototal := 0
	for i := 1; i < len(enum); i++ {
		if enum[i]%2 == 0 {
			etotal += enum[i]

		} else {
			ototal += enum[i]

		}
	}
	fmt.Println("The even number total is=", etotal)

	fmt.Println("The odd number total is=", ototal)

}

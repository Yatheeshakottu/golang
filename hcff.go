package main

import (
	"fmt"
)

func main() {
	x := 28
	y := 45
	hcf := 1
	for i := 1; i <= x && i <= y; i++ {
		if x%i == 0 && y%i == 0 {
			hcf = i
		}

	}
	fmt.Println(hcf)

}

package main

import (
	"fmt"
)

func main() {
	i := 0
	j := 10
	for ; i <= 10; i++ {
		fmt.Println(i, j)
		j--
	}
}

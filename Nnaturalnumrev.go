package main

import (
	"fmt"
)

func main() {
	var num, i int
	fmt.Print("Enter the natural number limit:")
	fmt.Scan(&num)
	for i = num; i >= 1; i-- {

		fmt.Println(i)
	}

}

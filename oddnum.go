package main

import (
	"fmt"
)

func main() {
	var i, num int
	fmt.Print("Enter the number")
	fmt.Scan(&num)
	for i = 1; i <= num; i++ {

		if i%2 != 0 {
			fmt.Println(i)
		}
	}

}

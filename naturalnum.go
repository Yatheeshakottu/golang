package main

import (
	"fmt"
)

func main() {
	var i, num int
	fmt.Print("Enter the N number")
	fmt.Scan(&num)
	for i = 1; i <= num; i++ {
		fmt.Println(i)
	}

}

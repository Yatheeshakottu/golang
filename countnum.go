package main

import (
	"fmt"
)

func main() {
	var num, count int
	fmt.Print("enter the number to count:")
	fmt.Scan(&num)
	for count = 0; num > 0; num = num / 10 {
		count++
	}
	fmt.Println("the total num of count:", count)

}

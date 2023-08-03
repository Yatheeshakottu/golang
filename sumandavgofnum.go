package main

import (
	"fmt"
)

func main() {
	var average, num, i int
	fmt.Print("Enter the number:")
	fmt.Scan(&num)
	total := 0
	for i = 1; i <= num; i++ {
		total = total + i
	}
	average = total / num
	fmt.Println(total)
	fmt.Println(average)
}

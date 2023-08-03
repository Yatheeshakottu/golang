package main

import (
	"fmt"
)

func main() {
	var num, sum int
	fmt.Print("enter the num")
	fmt.Scan(&num)
	sum = 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum = sum + i //sum+=i
		}
	}
	if sum == num {
		fmt.Println("it is a perfect number", num)
	} else {
		fmt.Println("it is not a perfect number", num)
	}
}

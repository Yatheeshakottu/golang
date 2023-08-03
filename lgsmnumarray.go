package main

import (
	"fmt"
)

func main() {
	var lgsm, minposition, maxposition int
	fmt.Print("enter the array size to find the largest=")
	fmt.Scan(&lgsm)
	lgsmArr := make([]int, lgsm)
	fmt.Print("enter the largest array items =")
	for i := 0; i < lgsm; i++ {
		fmt.Scan(&lgsmArr[i])
	}
	largest := lgsmArr[0]
	smallest := lgsmArr[0]
	for i := 0; i < lgsm; i++ {
		if largest < lgsmArr[i] {
			largest = lgsmArr[i]
			maxposition = i
		}
		if smallest > lgsmArr[i] {
			smallest = lgsmArr[i]
			minposition = i
		}
	}
	fmt.Println("\n the largest number in the lgsmArr=", largest)
	fmt.Println(" the index position of largest number =", maxposition)
	fmt.Println("\n the smallest number in the lgsmArr=", smallest)
	fmt.Println(" the index position of smallest number =", minposition)
}

package main

import (
	"fmt"
)

func main() {
	Avg := []int{4, 5, 6, 7, 2, 3, 9}
	fmt.Println(Avg)
	sum := 0
	for i := 0; i < len(Avg); i++ {
		sum += Avg[i]
	}
	arrayAvg := sum / len(Avg)
	fmt.Println("The average of array items", arrayAvg)
	fmt.Println("The sum of array items =", sum)

}

package main

import (
	"fmt"
	"math"
)

func main() {
	var sqrtnum, sqrtout float64
	//sqrtnum = 564277
	fmt.Print("Enter the number to find te square root:")
	fmt.Scan(&sqrtnum)
	sqrtout = math.Sqrt(sqrtnum)
	fmt.Println("the square root of the given number is:", sqrtout)

}

package main

import (
	"fmt"
	"math"
)

func main() {
	var pownum, expo float64
	fmt.Print("Enter the power of the number:")
	fmt.Scan(&pownum)
	fmt.Print("Enter the exponential :")
	fmt.Scan(&expo)
	power := math.Pow(pownum, expo)
	fmt.Println("The", pownum, "Of", expo, "IS", power)
}

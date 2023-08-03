package main

import (
	"fmt"
)

func main() {
	var pamount, sa, amount int
	fmt.Print("Enter the product amount:")
	fmt.Scan(&pamount)
	fmt.Print("Enter the sales amount:")
	fmt.Scan(&sa)
	if sa > pamount {
		amount = pamount - sa
		fmt.Println("the profit amount is:", amount)
	} else if pamount > sa {
		amount = pamount - sa
		fmt.Println("The loss amount is:", amount)

	} else {
		fmt.Println("There is no profit or loss")
	}
}

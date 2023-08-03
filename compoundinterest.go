package main

import (
	"fmt"
	"math"
)

func main() {
	var pinterest, timeperiod, cifuture, interestrate, compoundi float64
	fmt.Print("enter the principal interest")
	fmt.Scan(&pinterest)
	fmt.Print("enter the interest rate ")
	fmt.Scan(&interestrate)
	fmt.Print("enter the total no.of years")
	fmt.Scan(&timeperiod)
	cifuture = pinterest * (math.Pow((1 + interestrate/100), timeperiod))
	compoundi = cifuture - pinterest
	fmt.Println("The compound interest in future:", cifuture)
	fmt.Println("The compound interest :", compoundi)
}

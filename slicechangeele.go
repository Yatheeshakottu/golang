package main

import (
	"fmt"
)

func main() {
	x := []string{"Rainy", "winter", "summer"}
	fmt.Println(x)
	x[1] = "stroms"
	fmt.Println(x)
}

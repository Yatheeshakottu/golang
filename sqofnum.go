package main

import (
	"fmt"
)

func main() {
	var sqnum, sqrt int
	fmt.Print("Enter the square number")
	fmt.Scan(&sqnum)
	sqrt = sqnum * sqnum
	fmt.Println(sqrt)
}

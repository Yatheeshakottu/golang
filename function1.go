package main

import (
	"fmt"
)

func sum() (int, int) {
	return 1, 7
}
func main() {
	a, b := sum()//Here we use the 2 different return values from the call with multiple assignment.
	fmt.Println(a)
	fmt.Println(b)
	_, c := sum() // if  you want the subset of return values,we can use the blank identifier _
	fmt.Println(c)
}

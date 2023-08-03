package main

import "fmt"

func main() {
	var rows = 5
	var i, j int
	for i = rows; i > 0; i-- {
		for j = rows - i; j > 0; j-- {
			//fmt.Print(" 8")
			fmt.Printf("%d ", i)
			//fmt.Printf("%v ", j)
		}

		fmt.Println()
	}
}

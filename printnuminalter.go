package main

import "fmt"

func main() {
	rows := 7
	columns := 6
	for i := 1; i <= rows; i++ {
		for j := 1; j <= columns; j++ {
			if j%2 == 0 {
				fmt.Print("0 ")
			} else {
				fmt.Print("1 ")
			}
		}
		fmt.Println()
	}
}

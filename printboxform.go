package main

import "fmt"

func main() {
	rows := 7
	columns := 6
	for i := 1; i <= rows; i++ {
		for j := 1; j <= columns; j++ {
			if i == 1 || i == rows || j == 1 || j == columns {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}

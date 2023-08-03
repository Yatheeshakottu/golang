package main

import (
	"fmt"
)

func main() {
	var count, count1, k int
	var rows int = 6
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
			count++
		}
		for {
			if count <= rows-1 {
				fmt.Printf("%d", i+k)
				count++
			} else {
				count1++
				fmt.Printf("%d", i+k-2*count1)
			}
			k++
			if k == 2*i-1 {
				break
			}
		}
		k = 0
		count = 0
		count1 = 0
		fmt.Println("")
	}
}

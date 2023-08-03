package main

import (
	"fmt"
)

func main() {
	var count, count1, k int = 0, 0, 0
	for i := 1; i <= 5; i++ {
		k = 0
		for j := 1; j <= 5-i; j++ {
			fmt.Print(" ")
			count++//count=0
		}
		for {
			if k == 2*i-1 {
				break
			}
			if count <= 5-1 {
				fmt.Printf("%d ", i+k)
				//fmt.Printf("%d",i)
				count++
			} else {
				count1++
				//fmt.Printf("%d", i)
				fmt.Printf("%d ", i+k-2*count1)
				//	count1++
			}
			k++
		}
		count, count1, k = 0, 0, 0
		fmt.Println(" ")
	}
}

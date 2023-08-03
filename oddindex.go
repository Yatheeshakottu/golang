package main

import (
	"fmt"
)

func main() {
	num := []int{22, 3, 44, 55, 66}
	fmt.Println(num)
	for i := range num {//for i:=0;i<len(num):i++{
		if i%2 != 0 {
			fmt.Println(num[i])
		}
	}
}

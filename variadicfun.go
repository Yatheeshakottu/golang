package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println(nums, "")
	total := 0
	for _, c := range nums {
		total += c
	}
	fmt.Println(total)
}
func main() {
	sum(2, 3, 5, 6, 7)
	sum(3, 2, 6, 7, 8, 9, 122)
	x := []int{1, 2, 3, 4, 5, 6, 7}
	sum(x...)

}

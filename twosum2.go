package main

import (
	"fmt"
)

func twosum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := mp[diff]; ok {
			return []int{v, i}
		}
		mp[nums[i]] = i
	}
	return []int{-1, -1}
}
func main() {
	fmt.Println(twosum([]int{2, 7, 11, 17}, 9))

}

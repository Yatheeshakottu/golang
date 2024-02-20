package main

import (
	"fmt"
)

func twosum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

/*
	func twosummap(nums []int, target int) []int {
		mp := make(map[int]int)
		for i := 0; i < len(nums); i++ {
			diff := target - nums[i]
			if j, ok := mp[diff]; ok {
				return []int{j, i}
			}
			mp[nums[i]] = i
		}
		return []int{-1, -1}
	}
*/
func main() {
	fmt.Println(twosum([]int{2, 4, 5, 11, 15}, 9))
	//fmt.Println(twosummap([]int{2, 4, 5, 11, 15}, 9))
}
 
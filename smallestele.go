package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{12, 89, 34, 8, 567, 1, 789, 5}
	k := 2
	sort.Sort(sort.IntSlice(arr))
	for i := 0; i < k; i++ {
		fmt.Println(arr[i])
	}
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{34, 98, 46, 12, 908, 1234, 8, 1}
	k := 3
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	for i := 0; i < k; i++ {
		fmt.Println(arr[i])
	}

}

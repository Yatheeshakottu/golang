package main

import (
	"fmt"
)

func Duplicateinarray(arr []int) int {
	visited := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true {
			return arr[i]

		} else {
			visited[arr[i]] = true

		}
	}
	return -1
}
func main() {
	c := []int{1, 2, 3, 4, 1, 3, 6, 5, 2, 3}
	fmt.Println(Duplicateinarray(c))
	//fmt.Println(Duplicateinarray([]int{1, 5, 6, 7, 3, 1, 2, 2, 8, 5}))

}

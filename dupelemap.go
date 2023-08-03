package main

import "fmt"

func duplicate(s []int) {
	mp := make(map[int]int)
	for _, v := range s {
		mp[v]++
	}
	for _, v := range s {
		if mp[v] == 1 {
			fmt.Println("", v)
		}
	}
}
func main() {

	a := []int{1, 3, 2, 3, 5, 6, 7, 8, 10, 5, 4, 7}
	duplicate(a)
}

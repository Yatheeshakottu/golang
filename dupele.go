package main

import (
	"fmt"
)

func unic(a []int) []int {
	m := map[int]bool{}
	result := []int{}
	for v := range a {
		if m[a[v]] != true {
			m[a[v]] = true
			result = append(result, a[v])
		} 
	}
	return result
}
func main() {
	n := []int{2, 3, 5,6,4, 5, 7, 6, 9, 8, 1, 4}
	fmt.Println(unic(n))
}
         
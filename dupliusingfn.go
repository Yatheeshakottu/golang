package main

import "fmt"

func main() {
	x := []int{1, 3, 2, 5, 6, 6, 2, 7, 8, 9}
	duplicate(x)
}
func duplicate(x []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(x); i++ {
		if _, ok := m[x[i]]; ok {
			fmt.Println(x[i])
		}
		m[x[i]] = i
	}
	return x

}

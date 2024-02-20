package main

import "fmt"

func main() {
	x := []int{1, 3, 5, 6, 7, 5, 3, 8, 8}
	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i] == x[j] {
				fmt.Println(x[i])
			}
		}
	}
	fmt.Println("---------------")
	m := make(map[int]int)
	for i := 0; i < len(x); i++ {
		if _, ok := m[x[i]]; ok {
			fmt.Println(x[i])

		}

		m[x[i]] = i
	}
}

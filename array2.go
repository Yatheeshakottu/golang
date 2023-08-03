package main

import (
	"fmt"
	"sort"
)

func main() {
	m := []int{7, 6, 5, 1, 2}
	fmt.Println(m)
	sort.Ints(m)
	fmt.Println(m)
	i := 7
	for i = 0; i <= 7; i++ {
		i = i - 1
	}
	fmt.Println("The second largest num in the list", i)

}

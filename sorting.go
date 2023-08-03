package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"y", "a", "t", "h", "e", "e", "s", "h", "a"}
	fmt.Println(a)
	sort.Strings(a)
	//Sort(a)
	fmt.Println(a)
}

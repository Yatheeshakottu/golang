package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"h", "e", "l", "l", "o"} //Note that the sort is case-sensitive, so the uppercase letters will be ordered before lowercase letters in the output.
	fmt.Println(s)

	sort.Strings(s)
	fmt.Println(s)

}

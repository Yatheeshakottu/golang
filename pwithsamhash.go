package main

import (
	"fmt"
)

func main() {
	var country map[int]string
	country = make(map[int]string)
	country[1] = "india"
	country[2] = "america"
	country[3] = "london"
	country[4] = "china"
	for i, j := range country {
		fmt.Printf("key: %d value: %s\n", i, j)
	}
}

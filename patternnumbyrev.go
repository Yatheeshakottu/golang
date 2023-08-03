package main

import (
	"fmt"
)

func main() {
	num := 54321
	rev := 0
	for temp := num; temp > 0; temp = temp / 10 {
		remainder := temp % 10
		rev = rev*10 + remainder
		fmt.Println(rev)
	}
}

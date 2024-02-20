package main

import (
	"fmt"
)

func main() {
	sample := "Yathee"
	samplerune := []rune(sample)
	GeneratePermutation(samplerune, 0, len(samplerune)-1)
}
func GeneratePermutation(samplerune []rune, left, right int) {
	if left == right {
		fmt.Println(string(samplerune))
	} else {
		for i := left; i <= right; i++ {
			samplerune[left], samplerune[i] = samplerune[i], samplerune[left]
			GeneratePermutation(samplerune, left+1, right)
			samplerune[left], samplerune[i] = samplerune[i], samplerune[left]
		}

	}
}

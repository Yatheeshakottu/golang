package main

import (
	"fmt"
)

func LCM(n, m int) int {
	x := factors(n)
	y := factors(m)
	fmt.Println(x)
	fmt.Println(y)
	res := n * m
	for i := 1; i < len(x); i++ {
		for j := 1; j < len(y); j++ {
			if x[i] == y[j] {
				res = res / x[i]
				y = append(y[:j], y[j+1:]...)
			}
		}
	}
	return res

}
func factors(num int) []int {
	res := []int{}
	for i := 2; i <= num; i++ {
		for j := 2; j <= num; j++ {
			if num%i == 0 && num%j == 0 {
				res = append(res, i)
				num /= i
			}

		}
	}
	return res

}
func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Println(LCM(a, b))
}

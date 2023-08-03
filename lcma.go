package main

import (
	"fmt"
)

func LCM(n, m int) int {
	x := factors(n)
	y := factors(m)
	res := n * m // Initialize result as the product of n and m
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y); j++ {
			if x[i] == y[j] {
				res = res / x[i]              // Divide result by the common factor
				y = append(y[:j], y[j+1:]...) // Remove the common factor from y
				break
			}
		}
	}
	return res

}
func factors(num int) []int {
	res := []int{}
	for i := 2; i <= num; i++ {
		for num%i == 0 {
			res = append(res, i)
			num /= i
		}
	}
	return res

}
func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Println(LCM(a, b))
}

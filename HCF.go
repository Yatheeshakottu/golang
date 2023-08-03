package main

import "fmt"

func HCF(num1, num2 int) int {
	ar1 := GetFactors(num1)
	ar2 := GetFactors(num2)
	fmt.Println(ar1)
	fmt.Println(ar2)
	res := 1

	for i := len(ar1) - 1; i >= 0; i-- {
		for j := len(ar2) - 1; j >= 0; j-- {
			if ar1[i] == ar2[j] {
				res *= ar1[i]
				ar2 = remove(ar2, j)
				ar1 = remove(ar1, i)
				i = i - 1
				j = j - 1
			}
		}
	}
	return res
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func GetFactors(n int) []int {
	res := []int{}

	for i := 2; i <= n; {
		if n%i == 0 {
			res = append(res, i)
			n = n / i
			i = 2
		} else {
			i++
		}
	}
	return res
}

func main() {
	var w1, w2 int
	fmt.Scanln(&w1, &w2)

	fmt.Println(HCF(w1, w2))
}

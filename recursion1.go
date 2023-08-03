package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(6))
}
/*The fib() function executes as follows when called with an argument of 7:

1.fib(7) is called.
Since 7 is greater than or equal to 2, fib(7) returns fib(6) + fib(5).=13+8=21
2.fib(6) is called.
Since 6 is greater than or equal to 2, fib(6) returns fib(5) + fib(4).=8=5=13
3.fib(5) is called.
Since 5 is greater than or equal to 2, fib(5) returns fib(4) + fib(3).=5+3=8
4.fib(4) is called.
Since 4 is greater than or equal to 2, fib(4) returns fib(3) + fib(2).=3+2=5
5.fib(3) is called.
Since 3 is greater than or equal to 2, fib(3) returns fib(2) + fib(1).=2+1=3
6.fib(2) is called.
Since 2 is less than 2, fib(2) returns 2.=2
7.fib(1) is called.
Since 1 is less than 2, fib(1) returns 1.=1
fib(2) and fib(1) are added together and returned as the result of fib(3).
fib(4) now has the values of fib(3) and fib(2) and adds them together, returning 3.
 This process continues until fib(7) has the values of fib(6) and fib(5) and returns their sum, which is 13.*/




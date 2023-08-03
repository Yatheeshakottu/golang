package main

import (
	"fmt"
)

func main() {
	x := 1
	for {
		x++
		if x >= 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}

/*for n:=0;n<=5;n++{
	if n%2 !=0{
		continue
	}
	fmt.Println(n)
}
When n = 0, the if statement condition is false, so the value of n (i.e., 0) is printed.
When n = 1, the if statement condition is true (since 1%2 != 0), so the continue statement is executed, and the loop skips to the next iteration.
When n = 2, the if statement condition is false, so the value of n (i.e., 2) is printed.
When n = 3, the if statement condition is true (since 3%2 != 0), so the continue statement is executed, and the loop skips to the next iteration.
When n = 4, the if statement condition is false, so the value of n (i.e., 4) is printed.
When n = 5, the if statement condition is true (since 5%2 != 0), so the continue statement is executed, and the loop exits because n > 5.*/

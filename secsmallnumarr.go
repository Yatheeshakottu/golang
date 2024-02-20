package main

import (
	"fmt"
)

func main() {
	var small1, small2 int
	//x := []int{94, 23, 42, 432, 123, 4, 9, 2}
	x := []int{12, 3, 5, 6, 7, 9}
	if x[0] < x[1] {
		small1 = x[0]
		small2 = x[1]
	} else {
		small1 = x[1]
		small2 = x[0]
	}
	for i := 2; i < len(x); i++ {
		if x[i] < small1 {
			small2 = small1
			small1 = x[i]
		} else if x[i] < small2 && x[i] != small1 {
			small2 = x[i]
		}
	}
	/*if small1 == math.MaxInt {
		fmt.Println("All the elements in the array is same")
	} else {
		fmt.Println("The second smallest number in the array is :", small2)//if we print small1 we will ge smallest number in the array
	}*/
	fmt.Println("The second smallest number in the array is:", small2)

}

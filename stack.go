package main

import (
	"fmt"
)

// Stack follows a Last In First Out(LIFO)
func main() {
	var stack []string                        //[]int
	stack = append(stack, "yatheesha kottu ") //1,2,3,4//69891
	stack = append(stack, "kottu")            //6,67,7,3//78774
	for len(stack) > 0 {
		n := len(stack) - 1
		fmt.Println(stack[n])
		stack = stack[:n]
	}

}

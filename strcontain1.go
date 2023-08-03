/*Write a program which takes two strings as input from the user (str1 and str2). This program should print two strings as output (op1 and op2).

op1 should contain all the characters which are present in str1 but NOT present in str2.

op2 should contain all the characters which are present in str2 but NOT present in str1.



For example:


str1

str2

op1

op2

Example 1

ABC

BC

A

<null>

Example 2

BC

BANGALORE

C

ANGLORE*/

package main

import "fmt"

func main() {
	a := "ABCD"
	b := "ABE"
	op1 := ""
	op2 := ""
	m := make(map[rune]bool)
	for _, value := range a {
		m[value] = true
	}
	for _, key := range b {
		if _, ok := m[key]; !ok {
			op2 += string(key)
		} else {
			delete(m, key)
		}
	}
	for value := range m {
		op1 += string(value)
	}
	fmt.Println("OP1", op1)
	fmt.Println("op2", op2)

}

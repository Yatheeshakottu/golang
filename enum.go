package main

import "fmt"

func main() {
	var size, ecount, ocount int
	fmt.Println("enter the size of array")
	fmt.Scan(&size)
	dup := make([]int, size)
	fmt.Print("enter the array values")
	for i := 0; i < size; i++ {
		fmt.Scan(&dup[i])
	}
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			ecount++
		} else {
			ocount++
		}
	}
	fmt.Println("the even count", ecount)
	fmt.Println("The odd count:", ocount)
}

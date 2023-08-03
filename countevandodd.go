package main

import (
	"fmt"
)

func main() {
	var esum, osum, size int
	fmt.Print("Enter the array size")
	fmt.Scan(&size)
	num := make([]int, size)
	fmt.Print("Enter the number in to array:")
	for i := 0; i < size; i++ {
		fmt.Scan(&num[i])
	}
	esum = 0
	osum = 0
	for i := 0; i < size; i++ {
		if num[i]%2 == 0 {
			esum++
		} else {
			osum++
		}
	}
	fmt.Println(esum, osum)

}

//func main(){
//esum := 0
//osum := 0
//arr := []int{1, 2, 3, 4, 5, 6, 8, 7, 9, 10}
//	for i := 0; i < len(arr); i++ {
//if i%2 == 0 {
//esum++
//} else {
//osum++
//}
//}
//fmt.Println(esum)
//fmt.Println(osum)

//}

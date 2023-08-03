package main

import "fmt"

func main() {
	var reverse, palnum int
	fmt.Print("enter the number")
	fmt.Scanln(&palnum)
	reverse = 0
	for temp := palnum; temp > 0; temp = temp / 10 {
		remainder := temp % 10
		reverse = reverse*10 + remainder
	}
	fmt.Println("reverse of the given number:", reverse)

}

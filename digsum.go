package main

import (
	"fmt"
)

func main() {
	var digitsum, digitnum, digitremain int
	fmt.Print("Enter the digit number :")
	fmt.Scan(&digitnum)

	for digitsum = 0; digitnum > 0; digitnum = digitnum / 10 {

		digitremain = digitnum % 10
		digitsum = digitsum + digitremain
	}
	fmt.Println(digitsum)

}

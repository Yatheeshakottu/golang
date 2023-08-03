package main

import (
	"fmt"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result

	}
	return
}

//or
/*
	func reverse(s string)string{
		 rn:= []rune(s)
		for i,j:=0,len(s)-1;i<j;i,j=i+1,j-1{
			rn[i],rn[j]=rn[j],rn[i]
		}
	return string(rn)
		}
}
*/
func main() {
	str := "Yatheesha"
	strrev := reverse(str)
	fmt.Println(str)
	fmt.Println(strrev)

}

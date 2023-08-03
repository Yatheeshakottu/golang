package main

import (
	"golang.org/x/crypto/bcrypt"

	"fmt"
)

func main() {
	s := `password1234`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	loginPword1 := `password1234`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return

	}
	fmt.Println("YOU LOGGED IN")
}

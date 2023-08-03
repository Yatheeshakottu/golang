package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Last  string
	Age   int
}

func main() {
	u1 := user{
		First: "yatheesha",
		Last:  "kottu",
		Age:   21,
	}
	u2 := user{
		First: "james",
		Last:  "Bond",
		Age:   43,
	}
	u3 := user{
		First: "miss",
		Last:  "moneypenny",
		Age:   45,
	}
	users := []user{u1, u2, u3}
	fmt.Println(users)
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}

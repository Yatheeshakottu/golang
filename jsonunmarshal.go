package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":47},{"First":"Miss","Last":"Moneypenny","Age":37}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n all the data", people)
	for i, v := range people {
		fmt.Println("\n PERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}

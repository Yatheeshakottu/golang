package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json : "First"`
	Last  string `json :"Last"`
	Age   int    `json :"Age"`
}

func main() {
	s := `[{"First":"yatheesha","Last":"kottu","Age":21},{"First":"james","Last":"Bond","Age":43},{"First":"miss","Last":"moneypenny","Age":45}]`
	fmt.Println(s)
	var people []person//people:=[]person{}
	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range people {
		fmt.Println("\n all the data", i)
		fmt.Println(v.First, v.Last, v.Age)

	}
}

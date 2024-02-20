package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Place string `json:"place"`
}

func main() {
	//read the json file
	data, err := ioutil.ReadFile("yatheesha.json")
	if err != nil {
		log.Fatalf("error reading a json file %v ", err)
	}
	//parse json data into struct 
	var person Person
	err = json.Unmarshal(data, &person)
	if err != nil {
		log.Fatalf("error parsing json %v", err)
	}
	//print the parsed data
	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Println(person.Place)
}

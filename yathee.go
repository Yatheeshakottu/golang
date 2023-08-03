package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// define the JSON object
	jsonStr := `{
		"menu": {  
			"id": "file",  
			"value": "File",  
			"popup": {  
					"menuitem": [  
							{"value": "New", "onclick": "CreateDoc()"},  
							{"value": "Open", "onclick": "OpenDoc()"},  
							{"value": "Save", "onclick": "SaveDoc()"}  
					]  
			}  
		}
	}`

	// define a struct to hold the JSON data
	type MenuItem struct {
		Value   string `json:"value"`
		Onclick string `json:"onclick"`
	}

	type Popup struct {
		MenuItems []MenuItem `json:"menuitem"`
	}

	type Menu struct {
		ID    string `json:"id"`
		Value string `json:"value"`
		Popup Popup  `json:"popup"`
	}

	type Data struct {
		Menu Menu `json:"menu"`
	}

	// unmarshal the JSON data into the struct
	var data Data
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(err)
	}

	// print the JSON data
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}

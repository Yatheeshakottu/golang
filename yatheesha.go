package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := `{
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
	var m map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		panic(err)
	}
	jsonbytes, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonbytes))
}

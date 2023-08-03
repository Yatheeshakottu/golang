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

	var obj map[string]interface{}
	err := json.Unmarshal([]byte(data), &obj)
	if err != nil {
		panic(err)
	}

	// Print the JSON data
	jsonBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(jsonBytes))
}

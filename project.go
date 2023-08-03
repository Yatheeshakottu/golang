package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json: "Desc"`
	Content string `json:"Content"`
}

// let's declare global Article array
// that we can populate in our main function
// to simulate database
var Articles []Article

func homepage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "welcome to the homepage!")
	fmt.Println("Endpoint Hit: home page")

}
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	//DB
	json.NewEncoder(w).Encode(Articles)
}
func handleRequests() {
	http.HandleFunc("/", homepage)
	//add ur articles route and map it to our
	//returnAllArticles function like so
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article"},
	}
	handleRequests()
}

/*This code is a simple example of how to create a RESTful API in Go (Golang). It creates a server that listens on port 10000 and has two endpoints: "/" and "/articles".

The "main" function creates an "Articles" slice of "Article" structs, which are defined as having "Title", "Desc", and "Content" fields.

The "handleRequests" function uses the "http" package to handle incoming requests. The function "http.HandleFunc" maps a URL path to a function that will be called when that path is requested. In this case, the path "/" is mapped to the "homepage" function, and the path "/articles" is mapped to the "returnAllArticles" function.

The "homepage" function simply writes "welcome to the homepage!" to the client and logs a message to the console.

The "returnAllArticles" function logs a message to the console and encodes the global "Articles" slice as JSON and writes it to the client using the "json.NewEncoder" function.

Finally, the "main" function populates the "Articles" slice with two example article structs and calls the "handleRequests" function to start listening for incoming requests on port 10000.


*/

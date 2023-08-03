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

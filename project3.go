package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Income struct {
	Date     string `json:"date"`
	Account  string `json:"account"`
	Vendor   string `json:"vendor"`
	Category string `json:"Category"`
	Notes    string `json:"notes"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var (
	Articles []Article
	incomes  []Income
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	// DB
	json.NewEncoder(w).Encode(Articles)
}

func returnIncomeDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnIncomeDetails")
	// DB
	json.NewEncoder(w).Encode(incomes)
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	// add our articles route and map it to our
	// returnAllArticles function like so
	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/income", returnIncomeDetails)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	//file := io.Reader("c:\\")

	Articles = []Article{
		Article{Title: "Balaguruswamy C++", Desc: "C++ Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	incomes = []Income{
		Income{Date: time.Now().GoString(), Account: "Yetheesha", Vendor: "Bhagya", Category: "Dinner", Notes: "Party"},
		Income{Date: time.Now().GoString(), Account: "Bhagya", Vendor: "Yatheesha", Category: "Dinner", Notes: "Party"},
	}

	handleRequests()
}

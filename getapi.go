package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/get", getapi)
	http.HandleFunc("/geta", getapis)
	http.ListenAndServe(":8080", nil)
}
func getapi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func getapis(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "200")
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}



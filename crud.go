package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []movies

func getmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func deletemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().set("content-Type", "application/json")
	params := mux.vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)

			return

		}
	}
}
func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().set("content-Type", "application/json")
	var movie movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(item)

}
func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().set("content.Type", "application/json")
	params := mux.vars(r)
	for index, item := range movie {
		if item.ID == params["id"] {
			movies = append(movies, movie[:index], movie[index+1:]...)
			var movie movie
			_ = json.NewDecoder(r.Body).Decode((&movie))
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(r).Encode(movie)

		}
	}
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, movie{ID: "1", Isbn: "6478923", Title: "Movie one", Director: &Director{Firstname: "ss", Lastname: "Rajamouli"}})
	movies = append(movies, movie{ID: "2", Isbn: "754783", Title: "movie two", Director: &Director{Firstname: "Anil", Lastname: "ravipudi"}})
	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}	", getmovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE")
	r.HadleFunc("/movies", createmovie).Methods("POST")
	fmt.Println("Starting port at 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

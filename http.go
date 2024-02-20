package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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

var movies []movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	for _, item := range movies {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movieItem movie
	_ = json.NewDecoder(r.Body).Decode(&movieItem)
	movieItem.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movieItem)
	json.NewEncoder(w).Encode(movieItem)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	for index, item := range movies {
		if item.ID == id {
			var movieItem movie
			_ = json.NewDecoder(r.Body).Decode(&movieItem)
			movieItem.ID = id
			movies[index] = movieItem
			json.NewEncoder(w).Encode(movieItem)
			return
		}
	}
}

func main() {
	http.HandleFunc("/movies", getMovies)
	http.HandleFunc("/movies/{id}", createMovie).Methods("POST")
	http.HandleFunc("/movie", updateMovie).Methods("PUT")
	http.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	http.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	movies = append(movies, movie{ID: "1", Isbn: "6478923", Title: "Movie one", Director: &Director{Firstname: "ss", Lastname: "Rajamouli"}})
	movies = append(movies, movie{ID: "2", Isbn: "754783", Title: "movie two", Director: &Director{Firstname: "Anil", Lastname: "ravipudi"}})

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

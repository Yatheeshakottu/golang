package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User represents a simple user structure
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	// Set up the routes
	http.HandleFunc("/users", getUsersHandler)
	http.HandleFunc("/users/add", addUserHandler)
	http.HandleFunc("/users/delete", deleteUserHandler)

	// Start the server on port 8080
	fmt.Println("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// getUsersHandler handles GET requests to fetch all users
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// addUserHandler handles POST requests to add a new user
func addUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
}

// deleteUserHandler handles DELETE requests to remove a user
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var userToDelete User
	err := json.NewDecoder(r.Body).Decode(&userToDelete)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == userToDelete.ID {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

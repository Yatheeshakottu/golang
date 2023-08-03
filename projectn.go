package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to restfull API!")
}
func APIRequest() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
func main() {
	APIRequest()
}
/*This is a basic Go program that sets up a simple HTTP API that listens on port 3000 and responds to requests with a message.

The main function calls the apiRequests function, which sets up the API routes and starts the server.

The http.HandleFunc method is used to define a route for the root path ("/") of the API. The homePage function is called whenever a request is made to the root path. The homePage function writes a welcome message to the response body using the fmt.Fprintf method.

The log.Fatal method is used to log any errors that occur while the server is running. It is used here to log any errors that may occur when trying to listen on port 3000.

When the program is run, it will start a server that listens on port 3000. Whenever a request is made to the root path, the homePage function will be called and it will respond with the welcome message.*/

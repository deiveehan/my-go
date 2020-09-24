package main

import (
	"log"
	"net/http"
)

// Custom function that writes some content to a http Response Writer
func helloHandler(writer http.ResponseWriter, r *http.Request)  {
	writer.Write([]byte("Hello web user - from Deivee !!! "))
	log.Println(r.URL)
}


func main() {
  	// Delegating request to a handler
	http.HandleFunc("/hello", helloHandler)

	// Start a server and listen on port 8080
	err := http.ListenAndServe("localhost:8080", nil)

	// Display fatal error if the server has issues.
	log.Fatal(err)
}


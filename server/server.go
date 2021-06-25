package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// add handler function
	http.HandleFunc("/", handler)

	// specify address and start listening
	addr := "127.0.0.1:8080"
	log.Fatal(http.ListenAndServe(addr, nil))
}

// callback function which is called upon receiving a HTTP request
func handler(w http.ResponseWriter, r *http.Request) {
	message := "Welcome to Contact Server: %s"
	fmt.Fprintf(w, message, r.Method)
}

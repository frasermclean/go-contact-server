package main

import (
	"fmt"
	"log"
	"net/http"
)

// Data structure to store contact data
type Contact struct {
	Id           int
	FullName     string
	Email        string
	PhoneNumbers []string
}

func init() {
	c1 := Contact{
		Id:       1,
		FullName: "Bob Smith",
		Email:    "bob@acme.com",
	}
	c1.PhoneNumbers = append(c1.PhoneNumbers, "123")

	fmt.Printf("%+v\n", c1)
}

// Entry point for the program
func main() {
	// add handler function
	http.HandleFunc("/", handler)

	// specify address and start listening
	addr := "127.0.0.1:8080"
	log.Printf("Attempting to listen on: %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// Callback function which is called upon receiving a HTTP request
func handler(w http.ResponseWriter, r *http.Request) {
	message := "Welcome to Contact Server: %s"
	fmt.Fprintf(w, message, r.Method)
}

// Save the specfied contact to disk
func (c *Contact) save() error {
	return nil
}

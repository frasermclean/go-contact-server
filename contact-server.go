package main

import (
	"fmt"
	"io/ioutil"
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
	http.HandleFunc("/contact", contactHandler)

	// specify address and start listening
	addr := "127.0.0.1:8080"
	log.Printf("Attempting to listen on: %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// Callback function which is called upon receiving a HTTP request
func contactHandler(w http.ResponseWriter, r *http.Request) {
	var res string
	var err error

	// dispatch http method to correct function
	switch r.Method {
	case "GET":
		res, err = getContacts(w)

	case "POST":
		// try and read body
		body, ioError := ioutil.ReadAll(r.Body)
		if ioError != nil {
			log.Printf("Error reading body: %v", ioError)
			http.Error(w, "Error reading body.", http.StatusBadRequest)
			return
		}
		res, err = addContact(w, body)

	default:
		msg := "Unhandled method: " + r.Method
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// send response
	if err != nil {
		//
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, res)
}

func getContacts(w http.ResponseWriter) (string, error) {
	return "getContacts", nil
}

func addContact(w http.ResponseWriter, body []byte) (string, error) {
	return "addContact", nil
}

// Save the specified contact to disk
func (c *Contact) save() error {
	return nil
}

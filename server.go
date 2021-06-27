package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

	start := time.Now()

	// output logging message
	log.Println(fmt.Sprintf("Client %s request received from %s", r.Method, r.RemoteAddr))

	// dispatch http method to correct function
	switch r.Method {
	case "GET":
		res, err = getContacts()

	case "POST":
		// try and read body
		body, ioError := ioutil.ReadAll(r.Body)
		if ioError != nil {
			log.Printf("Error reading body: %v", ioError)
			http.Error(w, "Error reading body.", http.StatusBadRequest)
			return
		}
		err = addContact(body)

	default:
		msg := "Unhandled method: " + r.Method
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// check for error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set response body
	if res != "" {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, res)
	}

	message := fmt.Sprintf("Sent success response to %s, processing time: %s", r.RemoteAddr, time.Since(start))
	log.Println(message)
}

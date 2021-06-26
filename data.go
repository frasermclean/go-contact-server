package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// Data structure to store contact data
type Contact struct {
	Id           int
	FullName     string
	Email        string
	PhoneNumbers []string
}

var contacts []Contact

func init() {

	bob := Contact{
		Id:       1,
		FullName: "Bob Smith",
		Email:    "bob@acme.com",
	}
	bob.PhoneNumbers = append(bob.PhoneNumbers, "123")

	mary := Contact{
		Id:       2,
		FullName: "Mary Smith",
		Email:    "mary@acme.com",
	}
	mary.PhoneNumbers = append(mary.PhoneNumbers, "456")

	contacts = append(contacts, bob, mary)

	fmt.Printf("%+v\n", contacts)
}

// Returns a JSON serlialized string of all the current contacts
func getContacts() (string, error) {
	b, err := json.Marshal(contacts)
	if err != nil {
		log.Println(err)
		return "", err
	}

	s := string(b)
	return s, nil
}

// Add a new contact using the body as data values
func addContact(body []byte) (string, error) {
	fmt.Printf("body: %v\n", body)
	err := errors.New("not implemented")
	return "", err
}

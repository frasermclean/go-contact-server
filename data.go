package main

import (
	"encoding/json"
	"log"
)

// Data structure to store contact data
type Contact struct {
	Id           int      `json:"id"`
	FullName     string   `json:"full_name"`
	Email        string   `json:"email"`
	PhoneNumbers []string `json:"phone_numbers"`
}

var contacts []Contact

func init() {
	seedData()
}

// Adds sample contact data
func seedData() {
	bob := Contact{
		Id:       1,
		FullName: "Bob Smith",
		Email:    "bob@acme.com",
	}
	bob.PhoneNumbers = append(bob.PhoneNumbers, "0432 556 213")
	bob.PhoneNumbers = append(bob.PhoneNumbers, "03 3455 1235")

	mary := Contact{
		Id:       2,
		FullName: "Mary Smith",
		Email:    "mary@acme.com",
	}
	mary.PhoneNumbers = append(mary.PhoneNumbers, "0412 234 890")

	contacts = append(contacts, bob, mary)
}

// Returns a JSON serlialized string of all the current contacts
func getContacts() (string, error) {
	b, err := json.Marshal(contacts)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	s := string(b)
	return s, nil
}

// Add a new contact using the body as data values
func addContact(body []byte) error {

	// deserialize body into new contact
	contact := Contact{}
	err := json.Unmarshal(body, &contact)

	// chech for error
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// add new contact to existing contacts
	contacts = append(contacts, contact)
	log.Printf("Contact added: %v", contact)
	return nil
}

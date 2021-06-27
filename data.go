package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/mail"
	"regexp"
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
	jsonErr := json.Unmarshal(body, &contact)

	// chech for error
	if jsonErr != nil {
		log.Println(jsonErr.Error())
		return jsonErr
	}

	// validate new contact
	validErr := isValidContact(&contact)
	if validErr != nil {
		return validErr
	}

	// add new contact to existing contacts
	contacts = append(contacts, contact)
	log.Printf("Contact added: %v", contact)
	return nil
}

// Tests that the specified contact is valid
func isValidContact(contact *Contact) error {
	// test full name
	validFullName, _ := regexp.MatchString(`^[A-z ]+$`, contact.FullName)
	if !validFullName {
		return errors.New("error detected in full name: " + contact.FullName)
	}

	// test email
	_, mailErr := mail.ParseAddress(contact.Email)
	if mailErr != nil {
		return errors.New("error detected in email: " + contact.Email)
	}

	// test phone numbers
	for _, number := range contact.PhoneNumbers {
		validNumber, _ := regexp.MatchString(`^((\+61\s?)?(\((0|02|03|04|07|08)\))?)?\s?\d{1,4}\s?\d{1,4}\s?\d{0,4}$`, number)
		if !validNumber {
			return errors.New("error detected in phone number: " + number)
		}
	}

	return nil
}

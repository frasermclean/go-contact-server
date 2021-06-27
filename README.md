# Contact Server

Demo project written in Go with HTTP server that can manage contacts in a database.

## About

My first try at a project developed in the [Go language](https://golang.org).

## Running the code

Execute the following command from the project directory

```
go run .
```

## Usage

### Server Address

By default, the program starts a HTTP server running on localhost on post 8080. All contact operations should be directed to
the /contact relative path.

### Get Contacts

Request all contacts by sending a HTTP GET request to the server address. This will return all current contacts in a JSON array.

### Add New Contact

Send a HTTP POST request to the server address with JSON formatted contact data as body. Example below:

```
{
  "full_name": "Fraser McLean",
  "email": "example@email.com",
  "phone_numbers": [
    "400 123 456",
    "03 9213 1244"
  ]
}
```

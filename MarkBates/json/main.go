package main

import (
	"encoding/json"
	"os"
)

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:",omitempty"`
	Password  string `json:"-"`
}

func main() {

	/*
		u := User{
			FirstName: "Mark",
			LastName:  "Bates",
			Password:  "password",
		}
	*/

	u := User{}

	// encode the struct to json
	json.NewEncoder(os.Stdout).Encode(u)
}

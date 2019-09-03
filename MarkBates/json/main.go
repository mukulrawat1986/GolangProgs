package main

import (
	"encoding/json"
	"os"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string
	Password  string `json:"-"`
}

func main() {

	u := User{
		FirstName: "Mark",
		LastName:  "Bates",
		Password:  "password",
	}

	// encode the struct to json
	json.NewEncoder(os.Stdout).Encode(u)
}

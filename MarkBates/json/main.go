package main

import (
	"encoding/json"
	"log"
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

	// encode the struct to json
	// 	json.NewEncoder(os.Stdout).Encode(u)

	s := `
		{
			"first_name":"Mark",
			"LastName":"Bates"
		}
	`

	u := User{}

	// decode the json string into user object
	//	json.NewDecoder(strings.NewReader(s)).Decode(&u)

	json.Unmarshal([]byte(s), &u)

	log.Printf("u.FirstName: %s\n", u.FirstName)
	log.Printf("u.LastName: %s\n", u.LastName)
}

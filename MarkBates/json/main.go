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

func (u User) MarshalJSON() ([]byte, error) {

	m := struct {
		FirstName string
		LastName  string
	}{
		FirstName: "Mark",
		LastName:  "Bates",
	}

	return json.Marshal(m)
}
func main() {

	u := User{
		FirstName: "Mark",
		LastName:  "Bates",
	}

	// encode the struct to json
	json.NewEncoder(os.Stdout).Encode(u)

}

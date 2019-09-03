package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	Password  string
}

func main() {

	u := User{
		FirstName: "Mark",
		LastName:  "Bates",
		Password:  "password",
	}

	// encode the struct to json
	m, _ := json.Marshal(u)

	fmt.Println(string(m))
}

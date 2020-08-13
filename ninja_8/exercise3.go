package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name    string
	Age     int
	Sayings []string
}

func main() {
	user1 := User{
		Name:    "Ravi Shankar",
		Age:     30,
		Sayings: []string{"abc,xyz"},
	}
	user2 := User{
		Name:    "Sandhya",
		Age:     30,
		Sayings: []string{"abc,xyz"},
	}

	users := []User{user1, user2}
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

}

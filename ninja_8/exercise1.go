package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Age  int
}

func main() {
	user1 := user{
		Name: "Ravi Shankar",
		Age:  30,
	}
	user2 := user{
		Name: "Sandhya",
		Age:  27,
	}
	users := []user{user1, user2}
	fmt.Println(users)
	user_byte, _ := json.Marshal(users)
	fmt.Println(string(user_byte))

	u_byte, _ := json.Marshal(user1)
	fmt.Println(string(u_byte))
}

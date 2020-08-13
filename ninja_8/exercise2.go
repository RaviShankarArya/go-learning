package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:Name`
	Age  int    `json:Age`
}

func main() {
	str := `[{"Name":"Ravi Shankar","Age":30},{"Name":"Sandhya","Age":27}]`
	b_users := []byte(str)
	var usersStruct []User
	err := json.Unmarshal(b_users, &usersStruct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(usersStruct)
	for i, v := range usersStruct {
		fmt.Println("User", i+1, "Details")
		fmt.Println("Name :", v.Name)
		fmt.Println("Age :", v.Age)
	}

}

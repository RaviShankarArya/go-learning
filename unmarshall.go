package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"Name"`
	age  int    `json:"age"`
}

func main() {
	str := `[{"Name":"Ravi Shankar","Age":30},{"Name":"Sandhya","Age":27}]`
	byte_str := []byte(str)

	str1 := `{"Name":"Ravi Shankar","Age":30}`
	byte_str1 := []byte(str1)

	var people []person
	var people1 person
	json.Unmarshal(byte_str, &people)
	fmt.Println(people)

	json.Unmarshal(byte_str1, &people1)
	fmt.Println(people1)
}

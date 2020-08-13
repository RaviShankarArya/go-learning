package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := person{
		Name: "Ravi Shankar",
		Age:  30,
	}

	p2 := person{
		Name: "Sandhya",
		Age:  27,
	}
	people := []person{p1, p2}
	result, _ := json.Marshal(people)
	fmt.Println(string(result))

}

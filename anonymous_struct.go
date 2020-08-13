package main

import "fmt"

func main() {
	p1 := struct {
		first_name string
		last_name  string
		age        int
	}{
		first_name: "Ravi",
		last_name:  "Shankar",
		age:        30,
	}
	fmt.Println(p1)
}

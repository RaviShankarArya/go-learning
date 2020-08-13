package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int
}

func main() {
	p1 := person{first_name: "Ravi", last_name: "Shankar", age: 30}
	fmt.Println(p1)
	fmt.Println("First Name: ", p1.first_name)
	fmt.Println("Last Name: ", p1.last_name)
	fmt.Println("Age: ", p1.age)

}

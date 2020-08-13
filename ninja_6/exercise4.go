package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hello", p.first, p.last, "and your age is", p.age)
}

func main() {
	p := person{
		first: "Ravi",
		last:  "Shankar",
		age:   30,
	}

	fmt.Println(p)
	p.speak()
}

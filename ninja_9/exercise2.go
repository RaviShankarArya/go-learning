package main

import "fmt"

type person struct {
	name    string
	age     int
	sayings []string
}

type employee struct {
	person
	empId string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Name :", p.name, "Age :", p.age)
}

func (e *employee) speak() {
	fmt.Println("Name :", e.name, "Age :", e.age, "and empId :", e.empId)
}

func saySomething(h human) {
	fmt.Println(h)
	h.speak()
}

func main() {
	p := person{
		name:    "Ravi Shankar",
		age:     30,
		sayings: []string{"Ravi", "Shankru", "Arya"},
	}

	e := employee{
		person: p,
		empId:  "EMP001",
	}
	saySomething(&p)
	// saySomething(p)
	p.speak()

	saySomething(&e)
	e.speak()
}

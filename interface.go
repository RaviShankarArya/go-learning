package main

import "fmt"

type person struct {
	first_name string
	last_name  string
}

type employee struct {
	person
	is_active bool
}

type human interface {
	greeting()
}

func (emp employee) greeting() {
	fmt.Println("Welcome ,", emp.first_name)
}

func (p person) greeting() {
	fmt.Println("Welcome ,", p.first_name)
}

func display_full_name(h human) {
	fmt.Println("Interface")
	h.greeting()
	switch h.(type) {
	case person:
		fmt.Println("Full Name: ", h.(person).first_name, h.(person).last_name)
	case employee:
		fmt.Println("Full Name: ", h.(employee).first_name, h.(employee).last_name)
	}
}

func main() {
	emp := employee{
		person: person{
			first_name: "Ravi",
			last_name:  "Shankar",
		},
		is_active: true,
	}

	p := person{
		first_name: "Sandhya",
		last_name:  "Murthy",
	}

	fmt.Println(p)
	p.greeting()

	fmt.Println(emp)

	emp.greeting()

	display_full_name(emp)
	display_full_name(p)
}

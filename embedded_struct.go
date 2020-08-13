package main

import "fmt"

func main() {
	type person struct {
		first_name string
		last_name  string
		age        int
	}

	type employee struct {
		person
		first_name string
		emp_id     string
		is_active  bool
	}

	p1 := person{first_name: "Ravi", last_name: "Shankar", age: 30}
	fmt.Println(p1)
	fmt.Println("First Name: ", p1.first_name)
	fmt.Println("Last Name: ", p1.last_name)
	fmt.Println("Age: ", p1.age)

	emp := employee{
		person:     person{first_name: "Ravi", last_name: "Shankar", age: 30},
		first_name: "Ravi Shankar",
		emp_id:     "GGPL001",
		is_active:  true,
	}

	fmt.Println(emp)
	fmt.Println("EMP-ID: ", emp.emp_id)
	fmt.Println("Name: ", emp.first_name)
	fmt.Println("First Name: ", emp.person.first_name)
	fmt.Println("Last Name: ", emp.person.last_name)
	fmt.Println("Age: ", emp.person.age)
	fmt.Println("Is Actice?: ", emp.is_active)

}

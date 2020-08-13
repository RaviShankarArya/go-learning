package main

import "fmt"

type person struct{
	first_name string
	last_name string
}

type employee struct{
	person
	emp_id string
}

func main()  {
	emp1 := employee{
		person: person{
			first_name: "Ravi",
			last_name: "Shankar",
		},
		emp_id: "EMP001",
	}
	emp2 := employee{
		person: person{
			first_name: "Sandhya",
			last_name: "Murthy",
		},
		emp_id: "EMP002",
	}
	fmt.Println(emp1)
	emp1.greeting()
	emp2.greeting()
}

func (emp employee) greeting(){
	fmt.Println("Welcome", emp.first_name, emp.last_name,", your employee code is", emp.emp_id)
}
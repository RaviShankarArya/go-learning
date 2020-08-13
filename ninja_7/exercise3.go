package main

import "fmt"

type person struct {
	name string
}

func main() {
	p := person{
		name: "Ravi Shankar G",
	}
	fmt.Println(p.name)
	changeMe(&p)
}

func changeMe(p *person) {
	p.name = "Sandhya S Murthy"
	fmt.Println("Updated = ", p.name)
	fmt.Println("Updated = ", (*p).name)

}

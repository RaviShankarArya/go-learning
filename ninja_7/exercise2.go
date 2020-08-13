package main

import "fmt"

type person struct {
	name    string
	address string
}

func (p *person) changeMe() {
	p.address = "changed address"
	fmt.Println("Address =", p.address)
}

func main() {
	p := person{
		name:    "Ravi Shankar G",
		address: " current address",
	}
	fmt.Println(p)
	(&p).changeMe()
}

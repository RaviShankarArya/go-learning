package main

import "fmt"

type person struct {
	first_name          string
	last_name           string
	favorite_ice_creams []string
}

func main() {
	p1 := person{
		first_name:          "Ravi",
		last_name:           "Shankar",
		favorite_ice_creams: []string{"Choclate"},
	}

	p2 := person{
		first_name:          "Sandhya",
		last_name:           "Murthy",
		favorite_ice_creams: []string{"Strawberry"},
	}

	fmt.Println(p1)
	fmt.Println("First Name:", p1.first_name)
	fmt.Println("Last Name: ", p1.last_name)
	fmt.Println("Favorite Ice Creams: ")
	for i, v := range p1.favorite_ice_creams {
		fmt.Println(i+1, v)
	}
	fmt.Println(p2)
	fmt.Println("First Name:", p2.first_name)
	fmt.Println("Last Name: ", p2.last_name)
	fmt.Println("Favorite Ice Creams: ")
	for i, v := range p2.favorite_ice_creams {
		fmt.Println(i+1, v)
	}

}

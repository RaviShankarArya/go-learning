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

	persons := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}
	fmt.Println(persons)

	for k, v := range persons {
		fmt.Println(k)
		fmt.Println(v.first_name)
		for i, v := range v.favorite_ice_creams {
			fmt.Println(i, v)
		}
	}
}

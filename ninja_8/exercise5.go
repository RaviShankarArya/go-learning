package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name    string
	Age     int
	Sayings []string
}
type SortByName []User

func (a SortByName) Len() int           { return len(a) }
func (a SortByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type SortByAge []User

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	user1 := User{
		Name:    "Ravi Shankar",
		Age:     31,
		Sayings: []string{"abc", "xyz", "pqr"},
	}

	user2 := User{
		Name:    "Sandhya",
		Age:     27,
		Sayings: []string{"abc", "xyz", "pqr"},
	}

	user3 := User{
		Name:    "Rushant",
		Age:     6,
		Sayings: []string{"abc", "xyz", "pqr"},
	}

	users := []User{user1, user2, user3}
	fmt.Println(users)
	sort.Sort(SortByName(users))
	fmt.Println(users)
	sort.Sort(SortByAge(users))
	fmt.Println(users)

	for i, v := range users {
		fmt.Println("-----------------------")
		fmt.Println("User", i, "Details")
		fmt.Println("-----------------------")
		fmt.Println("Name :", v.Name)
		fmt.Println("Age :", v.Age)
		sort.Strings(v.Sayings)
		fmt.Println("Sayings :", v.Sayings)
		fmt.Println("-----------------------")

	}
}

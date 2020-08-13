package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hi")
	ai := []int{2, 7, 6, 8, 4, 5, 2, 3, 8}
	fmt.Println(ai)
	sort.Ints(ai)
	fmt.Println(ai)
	si := []string{"Ravi", "Sandhya", "Sanvi", "Rushanth", "Atharva"}
	fmt.Println(si)
	sort.Strings(si)
	fmt.Println(si)
}

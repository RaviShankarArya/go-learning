package main

import (
	"fmt"
	"sort"
)

func main() {
	age := []int{23, 13, 9, 43, 22, 90, 45}
	fmt.Println(age)
	sort.Ints(age)
	fmt.Println(age)
	name := []string{"Ravi", "Sandhya", "Rushant", "Roopa", "Sushant", "Ganesh", "Lakshmi"}
	fmt.Println(name)
	sort.Strings(name)
	fmt.Println(name)
}

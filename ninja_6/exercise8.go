package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	x := get_even_numbers(n...)
	fmt.Println(x())
}

func get_even_numbers(x ...int) func() []int {
	return func() []int {
		var even_numbers []int
		for _, v := range x {
			if v%2 == 0 {
				even_numbers = append(even_numbers, v)
			}
		}
		return even_numbers
	}
}

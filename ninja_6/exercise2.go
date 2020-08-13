package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := total(x...)
	fmt.Println(a)
	b := total1(x)
	fmt.Println(b)
}

func total(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func total1(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	even_number_sum := even(summation, x...)
	fmt.Println("Even Sum =", even_number_sum)

	odd_number_sum := odd(summation, x...)
	fmt.Println("Odd Sum =", odd_number_sum)

	highest := incrementor(highest_number, x...)
	fmt.Println(highest)

	x = []int{}
	highest = incrementor(highest_number, x...)
	fmt.Println(highest)
	x = []int{1}
	highest = incrementor(highest_number, x...)
	fmt.Println(highest)

}

func summation(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(x ...int) int, x ...int) int {
	var even []int
	for _, v := range x {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	return f(even...)
}

func odd(f func(x ...int) int, x ...int) int {
	var odd []int
	for _, v := range x {
		if v%2 != 0 {
			odd = append(odd, v)
		}
	}
	return f(odd...)
}

func incrementor(f func(x ...int) int, x ...int) int {
	n := f(x...)
	n++
	return n
}

func highest_number(x ...int) int {
	if len(x) == 0 {
		return 0
	}
	if len(x) == 1 {
		return 1
	}
	return x[0] + x[len(x)-1]
}

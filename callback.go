package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total := sum(x...)
	fmt.Println("Sum =", total)
	even_sum := even(sum, x...)
	fmt.Println("Even Numbers Sum =", even_sum)
	odd_sum := odd(sum, x...)
	fmt.Println("Odd Numbers Sum =", odd_sum)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(sum ...int) int, x ...int) int {
	var even []int
	for _, v := range x {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	return f(even...)
}

func odd(f func(sum ...int) int, x ...int) int {
	var odd []int
	for _, v := range x {
		if v%2 != 0 {
			odd = append(odd, v)
		}
	}
	return f(odd...)
}

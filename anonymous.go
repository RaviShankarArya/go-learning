package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := add(x...)
	fmt.Println("Sum =", sum)
	i := bar()
	fmt.Println(i())
}

func add(x ...int) int {
	return func(x ...int) int {
		sum := 0
		for _, v := range x {
			sum += v
		}
		return sum
	}(x...)
}

func bar() func() int {
	return func() int {
		return 451
	}
}

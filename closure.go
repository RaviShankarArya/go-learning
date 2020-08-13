package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println("a =", a())
	fmt.Println("a =", a())
	fmt.Println("a =", a())
	fmt.Println("a =", a())
	fmt.Println("b =", b())
	fmt.Println("b =", b())

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(x...)
	fmt.Println("sum =", s())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func sum(x ...int) func() int {
	return func() int {
		total := 0
		for _, v := range x {
			total += v
		}
		return total
	}
}

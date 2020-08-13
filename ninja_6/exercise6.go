package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := calculate(x...)
	fmt.Println("Sum =", sum)
}

func calculate(x ...int) int {
	return func() int {
		total := 0
		for _, v := range x {
			total += v
		}
		return total
	}()
}

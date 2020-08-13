package main

import (
	"fmt"
)

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := sum(n...)

	fmt.Println("Sum:", sum)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
		fmt.Println(v, sum)
	}
	return sum
}

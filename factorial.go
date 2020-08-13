package main

import "fmt"

func main() {
	n := 5
	x := factorial(n)
	fmt.Println("Factorial of ", n, "is", x)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

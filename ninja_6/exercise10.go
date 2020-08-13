package main

import "fmt"

func main() {
	a := increment()
	fmt.Println("a =", a())
	fmt.Println("a =", a())
	fmt.Println("a =", a())
	fmt.Println("a =", a())
	fmt.Println("a =", a())
	b := increment()
	fmt.Println("b =", b())
	fmt.Println("b =", b())
	fmt.Println("b =", b())
}

func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}

}

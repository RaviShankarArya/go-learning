package main

import "fmt"

func main() {
	a := 43
	b := &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)

	var x *int = b
	fmt.Println(x)
	fmt.Println(*x)

	a = 44
	fmt.Println(&a)
	fmt.Println(*&a)
	fmt.Println(a)
	*b = 22
	fmt.Println(*b)

}

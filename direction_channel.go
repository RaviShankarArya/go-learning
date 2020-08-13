package main

import "fmt"

func main() {
	c := make(chan int)
	cs := make(chan<- int)
	cr := make(<-chan int)

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cs)
	fmt.Printf("%T\n", cr)
}

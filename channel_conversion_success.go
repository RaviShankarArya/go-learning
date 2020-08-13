package main

import "fmt"

func main() {
	c := make(chan int)
	cs := make(chan<- int)
	cr := make(<-chan int)

	fmt.Printf("c : %T\n", c)
	fmt.Printf("cs : %T\n", cs)
	fmt.Printf("cr : %T\n", cr)

	fmt.Printf("c : %T\n", (chan<- int)(c))
	fmt.Printf("c : %T\n", (<-chan int)(c))
}

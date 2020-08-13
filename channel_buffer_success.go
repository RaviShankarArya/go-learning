package main

import "fmt"

func main() {
	c := make(chan int, 3)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
}

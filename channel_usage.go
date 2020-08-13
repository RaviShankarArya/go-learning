package main

import "fmt"

func main() {
	ch := make(chan int)

	go foo(ch)
	bar(ch)
	fmt.Println("about to exit")
}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}

package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()

	v, k := <-c

	fmt.Println(v, k)

	v, k = <-c

	fmt.Println(v, k)
}

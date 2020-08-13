package main

import (
	"fmt"
	"sync"
)

func main() {
	c := fanIn(boring("Ravi"), boring("Sandhya"))
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range c {
			fmt.Println(v)
			wg.Done()
		}

	}()
	wg.Wait()

	fmt.Println("about to exit")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		c <- msg
	}()
	return c
}

func fanIn(msg1, msg2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		c <- <-msg1
	}()
	go func() {
		c <- <-msg2
	}()
	return c
}

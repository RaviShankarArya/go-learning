package main

import (
	"fmt"
	"sync"
)

var gw sync.WaitGroup

func main() {
	fmt.Println("Hi")
	gw.Add(1)
	go foo()
	bar()
	gw.Wait()
}

func foo() {
	for i := 1; i <= 10; i++ {
		fmt.Println("foo =", i)
	}
	gw.Done()
}

func bar() {
	for i := 1; i <= 10; i++ {
		fmt.Println("bar =", i)
	}
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	fmt.Println("OS :", runtime.GOOS)
	fmt.Println("ARCH :", runtime.GOARCH)
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Go Routine :", runtime.NumGoroutine())
	waitGroup.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Go Routine :", runtime.NumGoroutine())
	waitGroup.Wait()
}

func foo() {
	for i := 1; i <= 10; i++ {
		fmt.Println("foo :", i)
	}
	waitGroup.Done()
}

func bar() {
	for i := 1; i <= 10; i++ {
		fmt.Println("bar :", i)
	}
}

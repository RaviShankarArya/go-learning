package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	// send
	go send(even, odd)
	// receive
	go receive(even, odd, fanIn)
	//fanIn
	for v := range fanIn {
		fmt.Println(v)
	}
	fmt.Println("about to close")
}

func send(even, odd chan<- int) {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(odd)
	close(even)
}

func receive(even, odd <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range even {
			fanIn <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanIn <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanIn)
}

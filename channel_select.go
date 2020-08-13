package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)
	// receive
	receive(even, odd, quit)
	fmt.Println("about to exit")
}

func send(e, o, q chan<- int) {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From Even Channel :", v)
		case v := <-o:
			fmt.Println("From Odd Channel :", v)
		case <-q:
			return
		}
	}
}

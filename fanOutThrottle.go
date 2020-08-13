package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
}

func populate(c chan<- int) {
	for i := 1; i <= 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1 <-chan int, c2 chan<- int) {
	var wg sync.WaitGroup
	const goroutines = 100
	wg.Add(goroutines)
	for i := 1; i <= goroutines; i++ {
		for v := range c1 {
			go func(n int) {
				c2 <- v
			}(v)

		}
		wg.Done()
	}
	wg.Wait()
	close(c2)
}

// func timeConsumption(n int) int {
// 	time.Sleep(time.Duration(rand.Intn(500)))
// 	return n + rand.Intn(500)
// }

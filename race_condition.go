package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	const gs = 100
	var waitGroup sync.WaitGroup
	waitGroup.Add(gs)
	// var mu sync.Mutex
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutines :", runtime.NumGoroutine())
	for i := 0; i < 100; i++ {
		go func() {
			// mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			// mu.Unlock()
			waitGroup.Done()
		}()
		fmt.Println("Goroutines :", runtime.NumGoroutine())
	}
	waitGroup.Wait()
	fmt.Println("Counter :", counter)
}

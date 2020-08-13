package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	const gs = 100
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(gs)
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutines :", runtime.NumGoroutine())
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Count :", counter)
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutines :", runtime.NumGoroutine())
}

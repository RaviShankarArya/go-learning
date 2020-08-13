package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutine :", runtime.NumGoroutine())
	for i := 0; i <= 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutine :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Count :", counter)
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutine :", runtime.NumGoroutine())
}

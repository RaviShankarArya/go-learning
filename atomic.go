package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutine :", runtime.NumGoroutine())
	var counter int64
	const gs = 100
	var waitGroup sync.WaitGroup
	waitGroup.Add(gs)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter =", atomic.LoadInt64(&counter))
			waitGroup.Done()
		}()
		fmt.Println("Goroutine :", runtime.NumGoroutine())
	}
	waitGroup.Wait()
	fmt.Println("Goroutine :", runtime.NumGoroutine())
	fmt.Println("Count =", counter)
}

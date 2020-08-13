package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	const gs = 10
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutine :", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < 10; i++ {
		func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Count :", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutine :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter :", counter)
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutine :", runtime.NumGoroutine())
}

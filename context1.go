package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Context Error check 1 :", ctx.Err())
	fmt.Println("Go routine :", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working :", n)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("Context Error check 2 :", ctx.Err())
	fmt.Println("Go routine :", runtime.NumGoroutine())

	fmt.Println("About to cancel")
	cancel()
	fmt.Println("Context Canceled")

	time.Sleep(time.Second * 2)
	fmt.Println("Context Error check 3 :", ctx.Err())
	fmt.Println("Go routine :", runtime.NumGoroutine())

}

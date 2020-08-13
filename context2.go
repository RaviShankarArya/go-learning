package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer fmt.Println("Cancelled Count")
	defer cancel()

	for v := range gen(ctx) {
		fmt.Println(v)
		if v == 10 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	c := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case c <- n:
				time.Sleep(time.Second * 3)
				n++
			}
		}
	}()
	return c
}

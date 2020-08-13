package main

import (
	"math/rand"
	"fmt"
	"sync"
	"time"
)

func main()  {
	c1:=make(chan int)
	c2:=make(chan int)
	go populate(c1)
	go fanInOut(c1,c2)

	for v:=range c2{
		fmt.Println(v)
	}
	
}

func populate(c chan int)  {
	for i:=1;i<=100;i++ {
		c<-i
	}
	close(c)
}

func fanInOut(c1,c2 chan int)  {
	fmt.Println("Hi")
	var wg sync.WaitGroup
	for v:=range c1{
		wg.Add(1)
		go func (v int)  {
			c2 <- timeConsumption(v)
			wg.Done()	
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumption(n int) int  {
	time.Sleep(time.Duration(rand.Intn(500)))
	return n + rand.Intn(500)
}
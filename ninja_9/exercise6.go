package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GO OS :", runtime.GOOS)
	fmt.Println("GOARCH ", runtime.GOARCH)
}

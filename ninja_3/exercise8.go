package main

import "fmt"

func main() {
	switch {
	case 2 == 2:
		fmt.Println("true")
	case 2 == 3:
		fmt.Println("false")
	default:
		fmt.Println("default")
	}
}

package main

import "fmt"

type ravi int

var x ravi

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}

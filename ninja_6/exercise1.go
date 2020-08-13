package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f)
	i, v := bar()
	fmt.Println(i, v)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "Hello"
}

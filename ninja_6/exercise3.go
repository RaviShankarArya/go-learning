package main

import "fmt"

func main() {
	defer foo()
	defer foo1()
	foo2()
	defer foo3()
	foo4()
}

func foo() {
	defer foo5()
	fmt.Println("foo")
}

func foo1() {
	fmt.Println("foo1")
}

func foo2() {
	fmt.Println("foo2")
}

func foo3() {
	fmt.Println("foo3")
}

func foo4() {
	fmt.Println("foo4")
}
func foo5() {
	fmt.Println("foo5")
}

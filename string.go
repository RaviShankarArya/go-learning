package main

import "fmt"

func main() {
	str := "Hello World!"
	fmt.Println(str)
	byte_str := []byte(str)
	fmt.Println(byte_str)
	fmt.Printf("%#U\n", byte_str)
	fmt.Printf("%#x\n", byte_str)
}

package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("Decimal - %d\n", x)
	fmt.Printf("Binary - %b\n", x)
	fmt.Printf("Hexa Decimal - %#x\n", x)
	a := x << 1

	fmt.Printf("Decimal - %d\n", a)
	fmt.Printf("Binary - %b\n", a)
	fmt.Printf("Hexa Decimal - %#x\n", a)

}

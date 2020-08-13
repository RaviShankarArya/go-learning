package main

import "fmt"

func main() {
	x := 1989
	for {
		if x > 2020 {
			break
		}

		fmt.Println(x)
		x++
	}
}

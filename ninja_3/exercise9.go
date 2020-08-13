package main

import "fmt"

func main() {
	switch x := "Skiing"; {
	case x == "Skiing":
		fmt.Println("case 1")
	case x == "Skiing":
		fmt.Println("case 1")
	default:
		fmt.Println("Default")
	}

}

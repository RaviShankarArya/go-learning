package main

import "fmt"

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(s)
	s1 := s[:3]
	fmt.Println(s1)
	s2 := s[6:]
	fmt.Println(s2)
	x := append(s1, s2...)
	fmt.Println(x)
	s = append(s[:3], s[6:]...)
	fmt.Println(s)

}

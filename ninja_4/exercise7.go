package main

import "fmt"

func main() {
	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	ms := [][]string{s1, s2}
	fmt.Println(ms)
	for i, ol := range ms {
		fmt.Println(i)
		for j, il := range ol {
			fmt.Println(j, il)
		}
	}
}

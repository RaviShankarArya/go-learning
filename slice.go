package main

import "fmt"

func main()  {
	x := []int{2,3,4,5}
	fmt.Println(x)
	y := []int{6,7,8,9}
	x = append(x,2,3,4)
	fmt.Println(x)
	z := append(x,y...)
	fmt.Println(z)

}
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Hi")
	s := `Password@1`
	bs, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	fmt.Println(s)
	s = string(bs)
	fmt.Println(bs)
	p := `Password@1`
	fmt.Println(s)
	err := bcrypt.CompareHashAndPassword(bs, []byte(p))
	if err != nil {
		fmt.Println("Password doesn't match")
		return
	}
	fmt.Println("You are logged In!!!")
}

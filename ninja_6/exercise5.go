package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type shape interface {
	area() float64
}

func info(s shape) {
	switch s.(type) {
	case circle:
		fmt.Println("area of circle =", s.(circle).area())
	case square:
		fmt.Println("area of square =", s.(square).area())
	}
	// fmt.Println(s.area())
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * 2)
}

func (s square) area() float64 {
	return s.length * 2
}

func main() {
	c := circle{
		radius: 200.0,
	}

	s := square{
		length: 230.5,
	}
	info(c)
	info(s)

}

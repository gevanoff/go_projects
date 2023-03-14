package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	tri := triangle{
		height: 2.0,
		base:   4.0,
	}

	squa := square{
		sideLength: 3.0,
	}

	printArea(tri)
	printArea(squa)
}

func printArea(s shape) {
	fmt.Println("Shape dimensions: ", s)
	area := float64(s.getArea())
	fmt.Println("Area equals:", area)
}

func (t triangle) getArea() float64 {
	area := float64(0.5 * t.base * t.height)
	return area
}

func (s square) getArea() float64 {
	area := (s.sideLength * s.sideLength)
	return area
}

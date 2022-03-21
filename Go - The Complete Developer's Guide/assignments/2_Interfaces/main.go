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
	s := square{sideLength: 5.5}
	printArea(s)

	t := triangle{base: 4.0, height: 6.0}
	printArea(t)
}

func printArea(s shape) {
	area := s.getArea()
	fmt.Println(area)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

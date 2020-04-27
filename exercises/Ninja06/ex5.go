package main

import "fmt"

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}
func (s circle) area() float64 {
	return s.radius * s.radius * 3.14
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq1 := square{
		side: 4,
	}
	ci := circle{
		radius: 4,
	}
	info(sq1)
	info(ci)
}

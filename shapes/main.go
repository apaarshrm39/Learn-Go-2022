package main

import (
	"fmt"
	"math"
)

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{
		side: 200,
	}

	tr := triangle{
		base:   100,
		height: 30,
	}

	printArea(tr)
	printArea(sq)
}

func (t triangle) getArea() float64 {
	return (0.5 * t.base * t.height)
}

func (s square) getArea() float64 {
	return (math.Pow(s.side, 2))
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

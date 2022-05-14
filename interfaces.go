package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (this rect) area() float64 {
	return this.width * this.height
}
func (this rect) perim() float64 {
	return 2*this.width + 2*this.height
}

func (this circle) area() float64 {
	return math.Pi * this.radius * this.radius
}
func (this circle) perim() float64 {
	return 2 * math.Pi * this.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

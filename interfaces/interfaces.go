package main

import (
	"fmt"
	"math"
)

type geometry interface { //interface for defining groups of methods
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) { //any variable with this interface can be call called with this method
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) { // can be used for checking the type of a variable holding an interface at runtime
	if c, ok := g.(circle); ok { //who variable defined here, then the ok is passed into the if statement
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 2}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}

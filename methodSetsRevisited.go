package main

import (
	"fmt"
	"math"
)

func main() {
	c := circle{5}

	fmt.Println("the area of the circle", c.area())
}

type shape interface {
	area() float64
}
type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func info(s shape) {
	fmt.Println("area", s.area())
}

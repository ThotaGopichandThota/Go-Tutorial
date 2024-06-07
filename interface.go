package main

import "fmt"

func main() {
	var length, width float64
	fmt.Println("enter the length of rectangle")
	fmt.Scanln(&length)
	fmt.Println("enter the width of rectangle")
	fmt.Scanln(&width)
	rectangle := Rectangle{Width: width, Length: length}
	details(rectangle)
	// circle := Circle{Radius: 6}
	// details(circle)
}

// interface with incomplete methods
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	Width  float64
	Length float64
}

// implementation of inter face for area method
func (r Rectangle) Area() float64 {

	area := r.Width * r.Length
	return area
}

// implementation of interface for peirmeter method
func (r Rectangle) Perimeter() float64 {

	perimenter := 2 * (r.Length + r.Width)
	return perimenter
}

// polymorphism method for interface
func details(s Shape) {
	fmt.Printf("Area: %f ,Perimeter :%f\n", s.Area(), s.Perimeter())
}

// type Circle struct {
// 	Radius float32
// }

// func (c Circle) Area() float32{
// 	area := 3.14 * c.Radius * c.Radius
// 	return area
// }
// func (c Circle) Perimeter() float32 {
// 	perimeter := 3.14 * c.Radius * 2
// 	return perimeter
// }

package structint

import "math"

// Shape is a general interface for all shapes having an Area method
type Shape interface {
	Area() float64
}

// Rectangle struct representing the rectangle shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle struct representing the circle shape
type Circle struct {
	Radius float64
}

// Area method calculates the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle represents the triangle shape
type Triangle struct {
	Base   float64
	Height float64
}

// Area method calculates the area of the triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

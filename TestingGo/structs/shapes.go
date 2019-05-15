package shapes

import "math"

// Shape interface represents all shapes with an area function
type Shape interface {
	Area() float64
}

// Rectangle type to represent a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle type to represent a circle
type Circle struct {
	Radius float64
}

// Triangle type to represent a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area method on Rectangle returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area method on Circle returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area method on Triangle returns the area of the triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

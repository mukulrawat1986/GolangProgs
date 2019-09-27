package structint

import "math"

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

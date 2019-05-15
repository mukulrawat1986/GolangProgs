package shapes

import "math"

// Rectangle type to represent a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle type to represent a circle
type Circle struct {
	Radius float64
}

// Area method on Rectangle returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area method on Circle returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

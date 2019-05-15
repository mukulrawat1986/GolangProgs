package shapes

// Rectangle type to represent a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of a rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

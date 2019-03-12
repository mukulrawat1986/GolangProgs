package StructsInterface

type Rectangle struct {
	Width  float64
	Height float64
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

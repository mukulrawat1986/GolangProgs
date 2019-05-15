package shapes

import "testing"

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  12.0,
			Height: 6.0,
		}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}

	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{
			Radius: 10.0,
		}

		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

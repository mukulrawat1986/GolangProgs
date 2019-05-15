package shapes

import "testing"

func TestArea(t *testing.T) {
	rectangle := Rectangle{
		Width:  12.0,
		Height: 6.0,
	}
	got := Area(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

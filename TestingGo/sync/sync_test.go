package _sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}

		// increment the counter 3 times
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("got %d want %d", counter.Value(), 3)
		}
	})

	t.Run("incrementing the counter 4 times and leaves it at 4", func(t *testing.T) {
		counter := Counter{}

		// increment the counter 4 times
		counter.Inc()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 4 {
			t.Errorf("got %d want %d", counter.Value(), 4)
		}
	})
}

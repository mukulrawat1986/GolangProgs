package _sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}

		// increment the counter 3 times
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter.Value(), 3)
	})

	t.Run("incrementing the counter 4 times and leaves it at 4", func(t *testing.T) {
		counter := Counter{}

		// increment the counter 4 times
		counter.Inc()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter.Value(), 4)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(wg *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter.Value(), wantedCount)
	})
}

func assertCounter(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

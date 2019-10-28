package c13sync

import (
	"sync"
	"testing"
)

func Test(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(counter, 3, t)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 10000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(wg *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}

		wg.Wait()

		assertCounter(counter, wantedCount, t)
	})
}

func assertCounter(counter *Counter, want int, t *testing.T) {
	if counter.Value() != want {
		t.Errorf("got %d want %d", counter.Value(), want)
	}
}

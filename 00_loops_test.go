package playing

import (
	"testing"
	"time"
)

func TestFors(t *testing.T) {
	t.Run("simple for loop", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			t.Log(i)
		}
	})

	t.Run("loop over slice/array", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for i, v := range s {
			t.Log(i, v)
		}
	})

	t.Run("loop over map", func(t *testing.T) {
		m := map[int]string{1: "a", 2: "b", 3: "c"}
		for k, v := range m {
			t.Logf("%d -> %s", k, v)
		}
	})

	t.Run("loop over channel", func(t *testing.T) {
		ch := make(chan int)
		go func() {
			for v := range ch {
				t.Log(v)
			}
		}()

		for i := 0; i < 10; i++ {
			ch <- i
		}
		time.Sleep(100 * time.Millisecond)
		close(ch)
	})

	// new in go 1.22
	t.Run("loop over int", func(t *testing.T) {
		for i := range 10 {
			t.Log(i)
		}

		for range 10 {
			t.Log("just a text")
		}
	})
}

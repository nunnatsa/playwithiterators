package mytree

import (
	"cmp"
	"iter"
	"maps"
	"math/rand/v2"
	"testing"
)

type collect[T cmp.Ordered] func(seq iter.Seq[T]) *Tree[T]

func BenchmarkCollect(b *testing.B) {
	for name, f := range map[string]collect[int]{
		"Collect":       Collect[int],
		"collectSlower": collectSlower[int],
	} {
		values := map[int]struct{}{}

		for len(values) < 1000 {
			values[rand.N[int](1000)] = struct{}{}
		}

		if len(values) != 1000 {
			b.Fatalf("we need 1000 values for this test, but got %d", len(values))
		}

		b.Run(name, func(b *testing.B) {
			for range b.N {
				l := f(maps.Keys(values))
				if ln := l.Len(); ln != len(values) {
					b.Errorf("should have %d elements, got %d", len(values), ln)
				}
			}
		})
	}
}

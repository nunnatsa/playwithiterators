package playing

import (
	"iter"
	"maps"
	"slices"
	"testing"
)

func TestPull(t *testing.T) {
	itr := slices.Values(numbers)
	next, stop := iter.Pull(itr)
	defer stop()

	for v, ok := next(); ok; v, ok = next() {
		t.Log(v)
	}
}

func TestPull2(t *testing.T) {
	itr := maps.All(numStr)
	next, stop := iter.Pull2(itr)
	defer stop()

	for k, v, ok := next(); ok; k, v, ok = next() {
		t.Logf("%#v: %#v\n", k, v)
	}
}

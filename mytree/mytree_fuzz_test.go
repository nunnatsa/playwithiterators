package mytree_test

import (
	"iter"
	"slices"
	"strings"
	"testing"

	"playing/mytree"
)

func FuzzCollectStrings(f *testing.F) {
	f.Add("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec luctus volutpat interdum. Maecenas nisl nisl, tincidunt cursus felis quis, ullamcorper aliquet tortor. Etiam ultrices neque iaculis libero sodales hendrerit. Fusce id lectus sapien. Donec porttitor ante ut mauris cursus, a ultrices arcu maximus. Pellentesque pretium, ipsum nec rutrum tristique, leo risus pulvinar orci, non tincidunt quam ante vitae lectus. Donec facilisis quam ac ligula sodales luctus.")
	f.Fuzz(func(t *testing.T, value string) {
		values := strings.Split(value, " ")
		t.Logf("values: %#v", values)
		tr := mytree.Collect(slices.Values(values))
		if len(values) < tr.Len() {
			t.Errorf("the tree lenght must be at least %d, but it's %d", len(values), tr.Len())
		}

		next, stop := iter.Pull(tr.Iter())
		defer stop()
		v, ok := next()
		if !ok {
			t.Errorf("should be at least one value")
		}
		last := v
		for v, ok = next(); ok; v, ok = next() {
			if v < last {
				t.Errorf("tree must be sorted")
			} else if v == last {
				t.Errorf("items in tree must be uniqu")
			}
			last = v
		}
	})
}

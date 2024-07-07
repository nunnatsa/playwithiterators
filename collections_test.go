package playing

import (
	"maps"
	"math/rand/v2"
	"slices"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"

	"playing/mylist"
	"playing/mytree"
)

func TestMyList(t *testing.T) {
	l := mylist.Collect(slices.Values(pods))

	if l.Len() != len(pods) {
		t.Errorf("list length should be %d, but it's %d", len(pods), l.Len())
	}

	flt := filter(l.Iter(), func(pod corev1.Pod) bool {
		_, ok := pod.Annotations["pick-me"]
		return !ok
	})

	trs := transform(flt, func(pod corev1.Pod) (types.UID, string) {
		return pod.UID, pod.Name
	})

	for uid, name := range trs {
		t.Logf(`{"uid"": %q, "name": %q}`, uid, name)
	}
}

func TestMyTree(t *testing.T) {
	tr := mytree.New[int]()

	length := 0
	for range 10 {
		if tr.Insert(rand.N[int](100)) {
			length++
		}
	}

	if tr.Len() != length {
		t.Errorf("tree size should be %d, but it's %d", length, tr.Len())
	}

	trs := transform(tr.Iter(), func(num int) (int, bool) {
		return num, num%2 == 0
	})

	rands := maps.Collect(trs)
	t.Logf("%#v\n", rands)
}

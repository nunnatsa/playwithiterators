package go1_23

import (
	"iter"
	"maps"
	"slices"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var pods = []corev1.Pod{
	{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pod-1",
			UID:  types.UID("11111111-1111-1111-1111-111111111111"),
		},
	},
	{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pod-2",
			UID:  types.UID("222222222-2222-2222-2222-22222222222"),
		},
	},
	{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pod-3",
			UID:  types.UID("33333333-3333-3333-3333-333333333333"),
		},
	},
	{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pod-4",
			UID:  types.UID("444444444-4444-4444-4444-44444444444"),
		},
	},
}

func transform[K, V1, V2 any](seq iter.Seq[V1], m func(V1) (K, V2)) iter.Seq2[K, V2] {
	return func(yield func(K, V2) bool) {
		for v1 := range seq {
			if !yield(m(v1)) {
				return
			}
		}
	}
}

func TestSliceToMap(t *testing.T) {
	itr := slices.Values(pods)
	podsByName := maps.Collect(transform(itr, func(pod corev1.Pod) (string, corev1.Pod) {
		return pod.Name, pod
	}))

	t.Logf("%#v\n", podsByName)

	nameByUID := maps.Collect(transform(itr, func(pod corev1.Pod) (types.UID, string) {
		return pod.UID, pod.Name
	}))
	t.Logf("%#v\n", nameByUID)
}

package playing

import (
	"maps"
	"slices"
	"testing"
)

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

var numStr = map[int]string{
	11: "11", 12: "12", 13: "13", 14: "14", 15: "15", 16: "16", 17: "17", 18: "18", 19: "19", 20: "20",
}

func TestSlicesAll(t *testing.T) {
	itr := slices.All(numbers)

	for idx, v := range itr { // same as: for idx, v := range numbers
		t.Log(idx, v)
	}
}

func TestSlicesValues(t *testing.T) {
	for v := range slices.Values(numbers) { // same as: for _, v := range numbers{}
		t.Log(v)
	}
}

func TestSlicesBackward(t *testing.T) {
	for idx, v := range slices.Backward(numbers) {
		t.Log(idx, v)
	}
}

//Now, maps:

func TestMapValues(t *testing.T) {
	itr := maps.Values(numStr)
	for v := range itr {
		t.Log(v)
	}
}

func TestSlicesCollect(t *testing.T) {
	itrV := maps.Values(numStr)
	sv := slices.Collect(itrV)
	t.Logf("%#v\n", sv)

	itrK := maps.Keys(numStr)
	sk := slices.Collect(itrK)
	t.Logf("%#v\n", sk)
}

func TestMapCollect(t *testing.T) {
	m := maps.Collect(slices.Backward(numbers))
	t.Logf("%#v\n", m)
}

func TestSlices(t *testing.T) {
	s := slices.AppendSeq(numbers, maps.Keys(numStr))
	t.Logf("%#v\n", s)

	slices.Sort(s)
	t.Logf("%#v\n", s)
}

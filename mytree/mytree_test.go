package mytree_test

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"

	"playing/mylist"
	"playing/mytree"
)

func TestTree_Insert(t *testing.T) {
	tr := mytree.New[int]()

	if !tr.Insert(5) {
		t.Error("insert 5 should return true")
	}
	if tr.Insert(5) {
		t.Error("insert 5 should return false")
	}
	if !tr.Insert(10) {
		t.Error("insert 10 should return true")
	}
	if !tr.Insert(12) {
		t.Error("insert 12 should return true")
	}
	if !tr.Insert(2) {
		t.Error("insert 2 should return true")
	}
	if tr.Insert(10) {
		t.Error("insert 10 should return false")
	}
	if !tr.Insert(3) {
		t.Error("insert 3 should return true")
	}
	if tr.Insert(5) {
		t.Error("insert 5 should return false")
	}
	if !tr.Insert(7) {
		t.Error("insert 7 should return true")
	}
	if !tr.Insert(-7) {
		t.Error("insert -7 should return true")
	}

	first := true
	var prev int
	for v := range tr.Values() {
		fmt.Println(v)
		if first {
			first = false
		} else {
			if v < prev {
				t.Errorf("%v should be greater than %v", v, prev)
			}
		}
		prev = v
	}

	s := slices.Collect(tr.Values())
	t.Log("tree before removing: \t" + tr.PrintDetails())
	for i, v := range s {
		tr.Remove(v)
		t.Logf("%v removed; tree now: \t%s", v, tr.PrintDetails())
		if l := tr.Len(); l != 7-(i+1) {
			t.Errorf("length should be %d, got %d", 7-(i+1), l)
		}
	}
}

func TestTree_Find(t *testing.T) {
	input := []string{"L", "R", "O", "U", "H", "J", "E"}
	tr := mytree.New[string]()

	for _, v := range input {
		if !tr.Insert(v) {
			t.Errorf("insert string for the first time should return true; string is %s", v)
		}
	}

	for _, v := range input {
		if !tr.Find(v) {
			t.Errorf("the string %q should be in the tree", v)
		}
	}

	for _, notFound := range []string{"", "A", "a string", "G"} {
		if tr.Find(notFound) {
			t.Errorf("string %q should not be in the tree", notFound)
		}
	}

	tr.Remove("H")
	if tr.Find("H") {
		t.Errorf(`"H" was removed for the tree abd should not be in it anymore`)
	}
}

func TestTree_Remove(t *testing.T) {
	for name, nums := range map[string][]int{
		"537": {5, 3, 7},
		"573": {5, 7, 3},
		"357": {3, 5, 7},
		"375": {3, 7, 5},
		"735": {7, 3, 5},
		"753": {7, 5, 3},
	} {
		t.Run(name, func(t *testing.T) {
			tr := getTree()
			fmt.Println(tr.PrintDetails())
			for i, v := range nums {
				tr.Remove(v)
				if l := tr.Len(); l != 7-(i+1) {
					t.Error("length should be ", 7-(i+1), " got ", l)
				}
			}
			fmt.Println(tr.PrintDetails())
		})
	}
}

func getTree() *mytree.Tree[int] {
	tree := mytree.New[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)
	return tree
}

func TestTree_Rand(t *testing.T) {
	for range 1000 {
		size := rand.N[int](100)
		t.Run(fmt.Sprintf("size=%d", size), func(t *testing.T) {
			expectedSize := size

			tr := mytree.New[int]()

			for range size {
				val := rand.N[int](1000)
				if rand.N[int](2)%2 == 1 {
					val *= -1
				}
				if !tr.Insert(val) {
					expectedSize--
				}
			}

			t.Logf("original size: %d; actual size: %d", size, expectedSize)

			if l := tr.Len(); l != expectedSize {
				t.Errorf("length should be %d, got %d", expectedSize, l)
			}

			vals := slices.Collect(tr.Values())
			t.Log(vals)
			if !slices.IsSorted(vals) {
				t.Errorf("values should be sorted")
			}

			if l := len(vals); l != expectedSize {
				t.Errorf("length should be %d, got %d", expectedSize, l)
			}

			list := mylist.New(vals...)
			for list.Len() > 0 {
				i := rand.N[int](list.Len())
				val, err := list.Remove(i)
				tr.Remove(val)
				if err != nil {
					t.Error(err)
				}
			}

			if l := tr.Len(); l != 0 {
				t.Errorf("length should be 0, got %d", tr.Len())
			}
		})
	}
}

func TestCollect(t *testing.T) {
	values := []int{5, 7, 10, 8, 6, 3, 2, 4}

	tree := mytree.Collect(slices.Values(values))
	if len(values) != tree.Len() {
		t.Errorf("tree length should be %d, but it's %d", len(values), tree.Len())
	}

	for v := range tree.Values() {
		t.Log(v)
	}
}

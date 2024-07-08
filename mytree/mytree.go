package mytree

import (
	"cmp"
	"fmt"
	"iter"
	"strings"
)

type node[T cmp.Ordered] struct {
	value T
	left  *node[T]
	right *node[T]
}

type Tree[T cmp.Ordered] struct {
	root *node[T]
}

// New creates a new Tree
func New[T cmp.Ordered]() *Tree[T] {
	return &Tree[T]{
		root: nil,
	}
}

func (n *node[T]) insert(value T) bool {
	if n.value == value {
		return false
	}

	if n.value > value {
		if n.right == nil {
			n.right = &node[T]{value: value}
			return true
		} else {
			return n.right.insert(value)
		}
	} else {
		if n.left == nil {
			n.left = &node[T]{value: value}
			return true
		} else {
			return n.left.insert(value)
		}
	}
}

func Collect[T cmp.Ordered](itr iter.Seq[T]) *Tree[T] {
	next, stop := iter.Pull(itr)
	defer stop()

	tr := New[T]()
	v, hasMore := next()
	if hasMore {
		tr.root = &node[T]{
			value: v,
		}

		for v, hasMore = next(); hasMore; v, hasMore = next() {
			tr.root.insert(v)
		}
	}

	return tr
}

func collectSlower[T cmp.Ordered](itr iter.Seq[T]) *Tree[T] {
	next, stop := iter.Pull(itr)
	defer stop()

	tr := New[T]()
	v, hasMore := next()
	if hasMore {
		tr.root = &node[T]{value: v}

		for v, hasMore = next(); hasMore; v, hasMore = next() {
			p := tr.root
			for {
				if p.value == v {
					break
				}

				if p.value > v {
					if p.right == nil {
						p.right = &node[T]{value: v}
						break
					}
					p = p.right

				} else {
					if p.left == nil {
						p.left = &node[T]{value: v}
						break
					}
					p = p.left
				}
			}
		}
	}

	return tr
}

func (t *Tree[T]) Insert(value T) bool {
	if t.root == nil {
		t.root = &node[T]{value: value}
		return true
	}

	return t.root.insert(value)
}

func (n *node[T]) traverse(yield func(T) bool) {
	if n == nil {
		return
	}

	if n.right != nil {
		n.right.traverse(yield)
	}
	yield(n.value)
	if n.left != nil {
		n.left.traverse(yield)
	}
}

func (t *Tree[T]) Iter() func(func(T) bool) {
	return func(yield func(T) bool) {
		for {
			t.root.traverse(yield)
			return
		}
	}
}

func (n *node[T]) find(value T) bool {
	if n.value == value {
		return true
	}
	if n.value > value {
		if n.right != nil {
			return n.right.find(value)
		}
	} else {
		if n.left != nil {
			return n.left.find(value)
		}
	}
	return false
}

func (t *Tree[T]) Find(value T) bool {
	if t.root == nil {
		return false
	}
	return t.root.find(value)
}

func (t *Tree[T]) Remove(value T) {
	if t.root != nil {
		t.root = t.root.remove(value)
	}
}

func (n *node[T]) remove(value T) *node[T] {
	if n == nil {
		return nil
	}

	if n.value > value {
		if n.right != nil {
			n.right = n.right.remove(value)
		}
	} else if n.value < value {
		if n.left != nil {
			n.left = n.left.remove(value)
		}
	} else {
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}

		minimum := n.value
		p := n.right
		for p != nil {
			minimum = p.value
			p = p.left
		}

		n.value = minimum
		if n.right != nil {
			n.right = n.right.remove(minimum)
		}
	}
	return n
}

func (t *Tree[T]) PrintDetails() string {
	return t.root.printDetails()
}

func (n *node[T]) printDetails() string {
	if n == nil {
		return "*"
	}
	b := strings.Builder{}
	b.WriteByte('[')
	b.WriteString(n.left.printDetails())
	b.WriteString(" <- ")
	b.WriteString(fmt.Sprintf("(%v)", n.value))
	b.WriteString(" -> ")
	b.WriteString(n.right.printDetails())
	b.WriteByte(']')

	return b.String()
}

func (t *Tree[T]) Len() int {
	return t.root.len()
}

func (n *node[T]) len() int {
	if n == nil {
		return 0
	}

	return n.left.len() + 1 + n.right.len()
}

package mylist

import (
	"fmt"
	"iter"
)

type item[T any] struct {
	value T
	next  *item[T]
}

type List[T any] struct {
	head   *item[T]
	length int
}

func New[T any](items ...T) *List[T] {
	l := &List[T]{}

	var p *item[T]
	for i, itm := range items {
		if i == 0 {
			l.head = &item[T]{value: itm}
			p = l.head
		} else {
			p.next = &item[T]{value: itm}
			p = p.next
		}
	}

	l.length = len(items)

	return l
}

func (l *List[T]) Len() int {
	return l.length
}

func (l *List[T]) Head() (T, error) {
	if l.head == nil {
		return *new(T), fmt.Errorf("trying to read from an empty list")
	}
	return l.head.value, nil
}

func (l *List[T]) Push(value T) {
	l.head = &item[T]{
		value: value,
		next:  l.head,
	}

	l.length++
}

func (l *List[T]) Pop() (T, error) {
	if l.head == nil {
		return *new(T), fmt.Errorf("trying to pop from an empty list")
	}

	v := l.head.value
	l.head = l.head.next

	l.length--

	return v, nil
}

func (l *List[T]) Insert(value T) {
	if l.head == nil {
		l.Push(value)
	} else {
		p := l.head
		for ; p.next != nil; p = p.next {
		}
		p.next = &item[T]{value: value}

		l.length++
	}
}

func (l *List[T]) Remove(index int) (T, error) {
	if l.head == nil {
		return *new(T), fmt.Errorf("trying to remove an empty list")
	}

	if index < 0 || index >= l.Len() {
		return *new(T), fmt.Errorf("index out of range")
	}

	if index == 0 {
		v := l.head.value
		l.head = l.head.next

		l.length--

		return v, nil
	}

	p := l.head
	var q *item[T]
	for ; index > 0; index-- {
		q = p
		p = p.next
	}
	q.next = p.next

	l.length--

	return p.value, nil
}

func (l *List[T]) Iter() func(func(T) bool) {
	return func(yield func(T) bool) {
		for p := l.head; p != nil; p = p.next {
			if !yield(p.value) {
				return
			}
		}
	}
}

func (l *List[T]) Iter2() func(func(int, T) bool) {
	return func(yield func(int, T) bool) {
		for i, p := 0, l.head; p != nil; i, p = i+1, p.next {
			if !yield(i, p.value) {
				return
			}
		}
	}
}

func Collect[T any](seq iter.Seq[T]) *List[T] {
	l := New[T]()

	next, stop := iter.Pull(seq)
	defer stop()

	v, ok := next()
	if ok {
		l.head = &item[T]{value: v}
		p := l.head

		l.length = 1

		for v, ok = next(); ok; v, ok = next() {
			p.next = &item[T]{value: v}
			p = p.next

			l.length++
		}
	}

	return l
}

func CollectSlower[T any](seq iter.Seq[T]) *List[T] {
	l := New[T]()
	for v := range seq {
		l.Push(v)
	}

	return l
}

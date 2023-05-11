// Package datastructuresgo defines common data structures.
package datastructuresgo

type node[T any] struct {
	data       T
	next, prev *node[T]
}

// A sorted doubly-linked list.
type SortedList[T any] struct {
	less       func(T, T) bool
	head, tail *node[T]
	count      uint
}

// NewSortedList, given a comparison function, returns a new SortedList.
func NewSortedList[T any](less func(T, T) bool) *SortedList[T] {
	return &SortedList[T]{
		less: less,
	}
}

// IsEmpty returns true when the SortedList has no nodes.
func (l SortedList[T]) IsEmpty() bool {
	return l.head == nil && l.tail == nil
}

// Add adds an item to the list.
func (l *SortedList[T]) Add(t T) {
	defer func() {
		l.count += 1
	}()

	n := &node[T]{
		data: t,
	}

	if l.IsEmpty() {
		l.head = n
		l.tail = n
		return
	}

	c := l.head
	for ; c != nil; c = c.next {
		if l.less(n.data, c.data) {
			n.next = c
			n.prev = c.prev
			c.prev = n
			if n.prev != nil {
				n.prev.next = n
			}
			if l.head == c {
				l.head = n
			} else if l.tail == c {
				l.tail = n
			}
			return
		}
	}

	// Couldn't find a larger element, so make it the new tail.
	l.tail.next = n
	n.prev = l.tail
	l.tail = n
}

// Count returns the number of nodes in the SortedList.
func (l SortedList[T]) Count() uint {
	return l.count
}

// Traverse traverses the nodes in the SortedList from head to tail. It calls
// the provided function for each, passing in the node's value. It stops either
// when the complete list is traversed, or the function returns false.
func (l SortedList[T]) Traverse(f func(T) bool) {
	for c := l.head; c != nil; c = c.next {
		if !f(c.data) {
			return
		}
	}
}

// Package datastructuresgo defines common data structures.
package datastructuresgo

// A single node in a sorted doubly-linked list.
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

// IsEmpty returns true when the list has no nodes.
func (l SortedList[T]) IsEmpty() bool {
	return l.head == nil && l.tail == nil
}

// AddHead adds an item to the list head.
func (l *SortedList[T]) AddHead(t T) {
	// TODO:: Implementation.
}

func (l SortedList[T]) Count() uint {
	return l.count
}

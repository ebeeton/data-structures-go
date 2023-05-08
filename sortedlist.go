// Package datastructuresgo defines common data structures.
package datastructuresgo

// A single Node in a sorted doubly-linked list.
type Node[T comparable] struct {
	Data       T
	Next, Prev *Node[T]
}

// A sorted doubly-linked list.
type SortedList[T comparable] struct {
	head, tail *Node[T]
}

// IsEmpty returns true when the list has no nodes.
func (l SortedList[T]) IsEmpty() bool {
	return l.head == nil && l.tail == nil
}

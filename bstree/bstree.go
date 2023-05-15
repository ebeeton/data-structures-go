// Package bstree provides a generic binary search tree.
package bstree

type node[T any] struct {
}

// A binary search tree.
type BSTree[T any] struct {
	less func(T, T) bool
	root *node[T]
}

// NewBSTree, given a comparison function, returns a new BSTree.
func NewBSTree[T any](less func(T, T) bool) *BSTree[T] {
	return &BSTree[T]{
		less: less,
	}
}

// IsEmpty returns true when there are no nodes in the tree.
func (t *BSTree[T]) IsEmpty() bool {
	return t.root == nil
}

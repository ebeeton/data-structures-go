// Package bstree provides a generic binary search tree.
package bstree

type node[T any] struct {
	data T
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
func (tree *BSTree[T]) IsEmpty() bool {
	return tree.root == nil
}

func (tree *BSTree[T]) Add(t T) {
	if tree.root == nil {
		tree.root = &node[T]{
			data: t,
		}
		return
	}
}

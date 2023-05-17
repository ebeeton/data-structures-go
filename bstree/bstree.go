// Package bstree provides a generic binary search tree.
package bstree

type node[T any] struct {
	data        T
	left, right *node[T]
}

// A binary search tree.
type BSTree[T any] struct {
	less  func(T, T) bool
	root  *node[T]
	count uint
}

// NewBSTree, given a comparison function, returns a new BSTree.
func NewBSTree[T any](less func(T, T) bool) *BSTree[T] {
	return &BSTree[T]{
		less: less,
	}
}

// IsEmpty returns true when there are no nodes in the BSTree.
func (tree *BSTree[T]) IsEmpty() bool {
	return tree.root == nil
}

// Add adds an item to the BSTree.
func (tree *BSTree[T]) Add(t T) {
	defer func() {
		tree.count += 1
	}()

	if tree.root == nil {
		tree.root = &node[T]{
			data: t,
		}
		return
	}

	n := tree.root
	for {
		if tree.less(t, n.data) {
			if n.left == nil {
				n.left = &node[T]{
					data: t,
				}
				return
			} else {
				n = n.left
			}
		} else {
			if n.right == nil {
				n.right = &node[T]{
					data: t,
				}
				return
			} else {
				n = n.right
			}
		}
	}
}

// Count returns the number of nodes in the BSTree.
func (tree *BSTree[T]) Count() uint {
	return tree.count
}

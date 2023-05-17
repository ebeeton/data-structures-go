// Package bstree provides a generic binary search tree.
package bstree

import "testing"

func less(a, b int) bool {
	return a < b
}

func TestIsEmpty(t *testing.T) {
	tree := NewBSTree(less)

	got := tree.IsEmpty()

	if !got {
		t.Errorf("IsEmpty() = %t, want true", got)
	}
}

func TestIsNotEmpty(t *testing.T) {
	tree := NewBSTree(less)
	tree.Add(10)

	got := tree.IsEmpty()

	if got {
		t.Errorf("IsEmpty() = %t, want false", got)
	}
}

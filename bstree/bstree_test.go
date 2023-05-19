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

func TestCount(t *testing.T) {
	tree := NewBSTree(less)
	tree.Add(10) // Root
	tree.Add(3)  // Root.left
	tree.Add(13) // Root.right
	tree.Add(1)  // Root.left.left
	tree.Add(15) // Root.right.right
	const want = 5

	got := tree.Count()

	if got != want {
		t.Errorf("Count() = %d, want %d", got, want)
	}
}

func TestPreOrder(t *testing.T) {
	tree := NewBSTree(less)
	tree.Add(10) // Root
	tree.Add(3)  // Root.left
	tree.Add(13) // Root.right
	tree.Add(1)  // Root.left.left
	tree.Add(15) // Root.right.right

	want := []int{10, 3, 1, 13, 15}

	idx := 0
	tree.PreOrder(func(i int) bool {
		if i != want[idx] {
			t.Errorf("PreOrder() = %d, want %d", i, want[idx])
		}
		idx += 1
		return true
	})
}

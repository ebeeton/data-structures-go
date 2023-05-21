// Package bstree provides a generic binary search tree.
package bstree

import "testing"

func less(a, b int) bool {
	return a < b
}

func equal(a, b int) bool {
	return a == b
}

func TestIsEmpty(t *testing.T) {
	tree := NewBSTree(less, equal)

	got := tree.IsEmpty()

	if !got {
		t.Errorf("IsEmpty() = %t, want true", got)
	}
}

func TestIsNotEmpty(t *testing.T) {
	tree := NewBSTree(less, equal)
	tree.Add(10)

	got := tree.IsEmpty()

	if got {
		t.Errorf("IsEmpty() = %t, want false", got)
	}
}

func TestCount(t *testing.T) {
	tree := NewBSTree(less, equal)
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
	tree := NewBSTree(less, equal)
	tree.Add(100)
	tree.Add(20)
	tree.Add(10)
	tree.Add(30)
	tree.Add(200)
	tree.Add(150)
	tree.Add(300)

	want := []int{100, 20, 10, 30, 200, 150, 300}

	idx := 0
	tree.PreOrder(func(i int) {
		if i != want[idx] {
			t.Errorf("PreOrder() = %d, want %d", i, want[idx])
		}
		idx += 1
	})
}

func TestInOrder(t *testing.T) {
	tree := NewBSTree(less, equal)
	tree.Add(100)
	tree.Add(20)
	tree.Add(10)
	tree.Add(30)
	tree.Add(200)
	tree.Add(150)
	tree.Add(300)

	want := []int{10, 20, 30, 100, 150, 200, 300}

	idx := 0
	tree.InOrder(func(i int) {
		if i != want[idx] {
			t.Errorf("InOrder() = %d, want %d", i, want[idx])
		}
		idx += 1
	})
}

func TestPostOrder(t *testing.T) {
	tree := NewBSTree(less, equal)
	tree.Add(100)
	tree.Add(20)
	tree.Add(10)
	tree.Add(30)
	tree.Add(200)
	tree.Add(150)
	tree.Add(300)

	want := []int{10, 30, 20, 150, 300, 200, 100}

	idx := 0
	tree.PostOrder(func(i int) {
		if i != want[idx] {
			t.Errorf("PostOrder() = %d, want %d", i, want[idx])
		}
		idx += 1
	})
}

func TestSearchPositive(t *testing.T) {
	tree := NewBSTree(less, equal)
	tree.Add(100)
	tree.Add(20)
	tree.Add(10)
	tree.Add(30)
	tree.Add(200)
	tree.Add(150)
	tree.Add(300)

	const want = 200

	if got := tree.Search(want); got == nil {
		t.Errorf("Search() = %d, want %d", got, want)
	}
}

func TestSearchNegative(t *testing.T) {
	tree := NewBSTree(less, equal)
	tree.Add(100)
	tree.Add(20)
	tree.Add(10)
	tree.Add(30)
	tree.Add(200)
	tree.Add(150)
	tree.Add(300)

	const want = 1

	if got := tree.Search(want); got != nil {
		t.Errorf("Search() = %d, want nil", got)
	}
}

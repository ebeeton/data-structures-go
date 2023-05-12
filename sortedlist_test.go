// Package datastructuresgo defines common data structures.
package datastructuresgo

import (
	"testing"
)

func less(a, b int) bool {
	return a < b
}

func TestIsEmpty(t *testing.T) {
	l := SortedList[int]{}

	got := l.IsEmpty()

	if !got {
		t.Errorf("IsEmpty() = %t, want true", got)
	}
}

func TestIsNotEmpty(t *testing.T) {
	l := NewSortedList(less)
	l.Add(1)

	got := l.IsEmpty()

	if got {
		t.Errorf("IsEmpty() = %t, want false", got)
	}
}

func TestNewSortedList(t *testing.T) {
	l := NewSortedList(less)
	count := l.Count()

	if l.less == nil {
		t.Errorf("l.less = nil, want less")
	} else if count != 0 {
		t.Errorf("Count() = %d, want 0", count)
	}
}

func TestAdd(t *testing.T) {
	l := NewSortedList(less)
	l.Add(2)
	l.Add(5)
	l.Add(3)
	l.Add(1)
	l.Add(4)
	l.Add(6)

	count := l.Count()

	if count != 6 {
		t.Errorf("Count() = %d, want 6", count)
	}

	want := []int{1, 2, 3, 4, 5, 6}
	idx := 0
	l.Traverse(func(i int) bool {
		if i != want[idx] {
			t.Errorf("Got %d, want %d", i, want[idx])
		}
		idx += 1
		return true
	})
}

func TestTraverseStopEarly(t *testing.T) {
	l := NewSortedList(less)
	l.Add(2)
	l.Add(5)
	l.Add(3)
	l.Add(1)
	l.Add(4)

	count := l.Count()

	if count != 5 {
		t.Errorf("Count() = %d, want 5", count)
	}

	idx := 0
	l.Traverse(func(i int) bool {
		idx += 1
		return idx != 3
	})

	if idx != 3 {
		t.Errorf("idx = %d, want 3", idx)
	}
}

func TestTraverseR(t *testing.T) {
	l := NewSortedList(less)
	l.Add(2)
	l.Add(5)
	l.Add(3)
	l.Add(1)
	l.Add(4)
	l.Add(6)

	count := l.Count()

	if count != 6 {
		t.Errorf("Count() = %d, want 5", count)
	}

	want := []int{6, 5, 4, 3, 2, 1}
	idx := 0
	l.TraverseR(func(i int) bool {
		if i != want[idx] {
			t.Errorf("Got %d, want %d", i, want[idx])
		}
		idx += 1
		return true
	})
}

func TestTraverseRStopEarly(t *testing.T) {
	l := NewSortedList(less)
	l.Add(2)
	l.Add(5)
	l.Add(3)
	l.Add(1)
	l.Add(4)

	count := l.Count()

	if count != 5 {
		t.Errorf("Count() = %d, want 5", count)
	}

	idx := 0
	l.TraverseR(func(i int) bool {
		idx += 1
		return idx != 3
	})

	if idx != 3 {
		t.Errorf("idx = %d, want 3", idx)
	}
}

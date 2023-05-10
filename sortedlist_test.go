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
	l.AddHead(1)

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

func TestAddHead(t *testing.T) {
	l := NewSortedList(less)
	l.AddHead(2)
	l.AddHead(5)
	l.AddHead(3)
	l.AddHead(1)
	l.AddHead(4)

	count := l.Count()

	if count != 5 {
		t.Errorf("Count() = %d, want 0", count)
	}
}

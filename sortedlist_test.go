// Package datastructuresgo defines common data structures.
package datastructuresgo

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	l := SortedList[int]{}

	got := l.IsEmpty()

	if !got {
		t.Errorf("IsEmpty() = %t, want true", got)
	}
}

func TestNewSortedList(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}

	l := NewSortedList(less)

	if l.less == nil {
		t.Errorf("l.less = nil, want less")
	}
}

// Package datastructuresgo defines common data structures.
package datastructuresgo

import "testing"

func TestIsEmpty(t *testing.T) {
	l := SortedList[int]{}

	got := l.IsEmpty()

	if !got {
		t.Errorf("IsEmpty() = %t, want true", got)
	}
}

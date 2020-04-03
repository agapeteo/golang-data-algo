package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	s := []int{3, 1, 4, 1, 9}
	MergeSort(s)

	exp := []int{1, 1, 3, 4, 9}
	for i, _ := range exp {
		if s[i] != exp[i] {
			t.Errorf("expected %v at index %v but got %v \n actual slice: %v", exp[i], i, s[i], s)
		}
	}
}

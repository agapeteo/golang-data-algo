package list

import (
	"testing"
)

func TestSimpleLinkedList_Sort(t *testing.T) {
	list := SimpleLinkedList{}
	list.Append(3)
	list.Append(1)
	list.Append(2)
	list.Append(2)
	list.Append(5)

	list.SortInts()

	elements := list.Elements()
	expected := []interface{}{1, 2, 2, 3, 5}
	if !isEqual(expected, elements) {
		t.Errorf("expecting %v, but got %v", expected, elements)
	}
}

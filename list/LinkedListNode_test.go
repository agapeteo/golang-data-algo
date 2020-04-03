package list

import (
	"testing"
)

func TestNode_Elements(t *testing.T) {
	node := CreateLinkedListNode(0)

	node.Append(1)
	node.Append(2)
	node.Append(3)

	elements := node.Elements()
	//fmt.Printf("elements: %v, length: %v ", elements, node.Size())
	if !isEqual([]interface{}{0, 1, 2,}, elements) {
		t.Errorf("expected equal slices, but got %v", elements)
	}
}

func isEqualInterfaceSlices(left, right []interface{}) bool {
	if left == nil || right == nil {
		return false
	}
	for i, _ := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func TestNode_RemoveAt(t *testing.T) {
	// given
	node := CreateLinkedListNode(0)

	node.Append(1)
	node.Append(2)
	node.Append(3)
	node.Append(4)

	// when 
	node.RemoveAt(3)

	// then
	expected := []interface{}{0, 1, 2, 4}
	actual := node.Elements()
	if !isEqualInterfaceSlices(actual, expected) {
		t.Errorf("expected equal slices, but got %v", actual)
	}
}

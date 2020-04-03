package list

import (
	"fmt"
	"testing"
)

func TestLinkedListNode_Deduplicate(t *testing.T) {
	node := CreateLinkedListNode(0)

	node.Append(1)
	node.Append(2)
	node.Append(2)
	node.Append(3)
	node.Append(3)
	node.Append(2)
	node.Append(4)
	node.Deduplicate()

	elements := node.Elements()
	fmt.Printf("elements: %v, length: %v ", elements, node.Size())
	if !isEqual([]interface{}{0, 1, 2, 3, 4}, elements) {
		t.Errorf("expected equal slices, but got %v", elements)
	}
}
package tree

import (
	"testing"
)

func TestMaxHeap_Push_Peak_Size(t *testing.T) {
	// given
	var h Heap = &MaxHeap{}
	var example = []int{5, 10, 2, -1, 0}

	// when
	for i := 0; i < len(example); i++ {
		h.Push(example[i])
	}

	// then
	var actual, _ = h.Peak()
	if actual != 10 {
		t.Errorf("expected 10 but got %v", actual)
	}

	if h.Size() != len(example) {
		t.Errorf("expected %v but got %v", len(example), h.Size())
	}
	if h.Size() != len(h.Elements()) {
		t.Errorf("expected %v but got %v, elements: %v", h.Size(), len(h.Elements()), h.Elements())
	}
}

func TestMaxHeap_Pop(t *testing.T) {
	// given
	var h Heap = &MaxHeap{}
	var example = []int{5, 10, 2, -1, 0}

	// when
	for i := 0; i < len(example); i++ {
		h.Push(example[i])
	}

	// then
	var actual = make([]int, 0)

	for h.Size() > 0 {
		var v, _ = h.Pop()
		actual = append(actual, v)
	}

	var expected = []int{10, 5, 2, 0, -1}

	if !isEqual(actual, expected) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

package list

import (
	"testing"
)

func TestMinStack_Min(t *testing.T) {
	// given
	stack := CreateMinStack()

	// when
	stack.Push(1)
	stack.Push(2)
	stack.Push(0)
	stack.Push(0)
	stack.Push(-1)
	stack.Push(3)

	// then
	if stack.Min() != -1 {
		t.Errorf("expected -1, but got %v", stack.Min())
	}
	stack.Pop()
	stack.Pop()

	if stack.Min() != 0 {
		t.Errorf("expected 0, but got %v", stack.Min())
	}
	stack.Pop()

	if stack.Min() != 0 {
		t.Errorf("expected 0, but got %v", stack.Min())
	}
	stack.Pop()

	if stack.Min() != 1 {
		t.Errorf("expected 1, but got %v", stack.Min())
	}
}
package list

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	// given
	// when
	var list LinkedList = &SimpleLinkedList{}

	// then
	if list.Size() != 0 {
		t.Errorf("list size should be empty, but got %v", list.Size())
	}
	if list.GetFirst() != nil {
		t.Errorf("list head should be nil")
	}
	if list.GetLast() != nil {
		t.Errorf("list tail should be nil")
	}
}

func TestSimpleLinkedList_Get_empty(t *testing.T) {
	// given
	// when
	var list LinkedList = &SimpleLinkedList{}

	// then
	var _, actualError = list.Get(0)

	if actualError == nil {
		t.Errorf("error should not be nil")
	}

	if actualError != EmptyListError {
		t.Errorf("expected error with type EmptyListError but got %v ", actualError)
	}
}

func TestSimpleLinkedList_Get_outOfBounds(t *testing.T) {
	// given
	// when
	var list LinkedList = &SimpleLinkedList{}
	list.Append("something")

	// then
	var _, actualError = list.Get(1)

	if actualError == nil {
		t.Errorf("error should not be nil")
	}

	if actualError != IndexOutOfBoundsError {
		t.Errorf("expected error with type IndexOutOfBoundsError but got %v ", actualError)
	}
}

func TestSimpleLinkedList_GetFirst_GetLast_Get(t *testing.T) {
	// given
	var list LinkedList = &SimpleLinkedList{}
	const first = 1
	const second = 2
	const last = 3

	// when
	list.Append(second)
	list.Append(last)
	list.AddFirst(first)

	// then
	const expectedSize = 3
	if list.Size() != expectedSize {
		t.Errorf("list size should be 3 but got %v", list.Size())
	}

	if list.GetFirst() != first {
		t.Errorf("first value should be %v, but got %v", first, list.GetFirst())
	}

	if list.GetLast() != last {
		t.Errorf("first value should be %v, but got %v", first, list.GetFirst())
	}

	actualSecond, _ := list.Get(1)
	if actualSecond != second {
		t.Errorf("second value should be %v but got %v", second, actualSecond)
	}
}

func TestSimpleLinkedList_Append(t *testing.T) {
	// given
	var list LinkedList = &SimpleLinkedList{}

	// when
	list.Append(1)
	list.Append(2)

	// then
	expected := []interface{}{1, 2}
	if !isEqual(expected, list.Elements()) {
		t.Errorf("elements not equal. expected %v but got: %v", expected, list.Elements())
	}
}

func TestSimpleLinkedList_AddFirst(t *testing.T) {
	// given
	var list LinkedList = &SimpleLinkedList{}

	// when
	list.AddFirst(1)
	list.AddFirst(2)
	list.AddFirst(3)

	// then
	expected := []interface{}{3, 2, 1}
	if !isEqual(expected, list.Elements()) {
		t.Errorf("elements not equal. expected %v but got: %v", expected, list.Elements())
	}
}

func TestSimpleLinkedList_DeleteFirst(t *testing.T) {
	// given
	var list LinkedList = &SimpleLinkedList{}

	// when
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.DeleteFirst()
	list.DeleteFirst()

	// then
	expected := []interface{}{3}
	if !isEqual(expected, list.Elements()) {
		t.Errorf("elements not equal. expected %v but got: %v", expected, list.Elements())
	}
}

func TestSimpleLinkedList_DeleteLast(t *testing.T) {
	// given
	var list LinkedList = &SimpleLinkedList{}

	// when
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.DeleteLast()
	list.DeleteLast()

	// then
	expected := []interface{}{1}
	if !isEqual(expected, list.Elements()) {
		t.Errorf("elements not equal. expected %v but got: %v", expected, list.Elements())
	}
}

func TestSimpleLinkedList_Elements(t *testing.T) {
	// given
	var list LinkedList = &SimpleLinkedList{}

	// when
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// then
	expected := []interface{}{1, 2, 3}
	if !isEqual(expected, list.Elements()) {
		t.Errorf("elements not equal. expected %v but got: %v", expected, list.Elements())
	}
}

func TestSimpleLinkedList_ReverseElements(t *testing.T) {
	// given
	var list LinkedList = &SimpleLinkedList{}

	// when
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// then
	expected := []interface{}{3, 2, 1}
	if !isEqual(expected, list.ReverseElements()) {
		t.Errorf("elements not equal. expected %v but got: %v", expected, list.ReverseElements())
	}
}

func TestSimpleLinkedList_Elements_empty(t *testing.T) {
	// given
	// when
	var list LinkedList = &SimpleLinkedList{}

	// then
	expected := make([]interface{}, 0)
	if !isEqual(expected, list.Elements()) {
		t.Errorf("elements not equal. expected %v but got: %v", expected, list.Elements())
	}
}

func TestSimpleLinkedList_FirstIndexOf(t *testing.T) {
	// given
	var list LinkedList = &SimpleLinkedList{}

	// when
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// then
	actual := list.FirstIndexOf(3)
	const expected = 2
	if actual != expected {
		t.Errorf("first index should be %v but got %v", expected, actual)
	}
}

func TestSimpleLinkedList_FirstIndexOf_missing(t *testing.T) {
	// given
	var list LinkedList = &SimpleLinkedList{}

	// when
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// then
	actual := list.FirstIndexOf(5)
	const expected = -1
	if actual != expected {
		t.Errorf("first index should be %v but got %v", expected, actual)
	}
}

func TestSimpleLinkedList_FirstIndexOf_empty(t *testing.T) {
	// given
	// when
	var list LinkedList = &SimpleLinkedList{}

	// then
	actual := list.FirstIndexOf(5)
	const expected = -1
	if actual != expected {
		t.Errorf("first index should be %v but got %v", expected, actual)
	}
}

func TestSimpleLinkedList_InsertAfterIndex(t *testing.T) {
	// given
	var list LinkedList = &SimpleLinkedList{}

	// when
	list.Append(1)
	list.Append(3)

	e := list.InsertAfterIndex(0, 2)
	if e != nil {
		t.Errorf("error happend durint insersion: %v", e)
	}

	// then
	expected := []interface{}{1, 2, 3}
	if !isEqual(expected, list.Elements()) {
		t.Errorf("elements not equal. expected %v but got: %v", expected, list.Elements())
	}
}

func TestSimpleLinkedList_Reverse(t *testing.T) {
	// given
	var list = SimpleLinkedList{}

	// when
	list.Append(3)
	list.Append(2)
	list.Append(1)

	list.Reverse()

	// then
	expected := []interface{}{1, 2, 3}
	if !isEqual(expected, list.Elements()) {
		t.Errorf("elements not equal. expected %v but got: %v", expected, list.Elements())
	}
}

func TestSimpleLinkedList_Empty(t *testing.T) {
	// given
	var list = SimpleLinkedList{}

	// when

	list.Reverse()

	// then
	var expected = make([]interface{}, 0)
	if !isEqual(expected, list.Elements()) {
		t.Errorf("elements not equal. expected %v but got: %v", expected, list.Elements())
	}
}

func TestSimpleLinkedList_TwoElements(t *testing.T) {
	// given
	var list = SimpleLinkedList{}

	// when
	list.Append(2)
	list.Append(1)

	list.Reverse()

	// then
	expected := []interface{}{1, 2}
	if !isEqual(expected, list.Elements()) {
		t.Errorf("elements not equal. expected %v but got: %v", expected, list.Elements())
	}
}

func isEqual(a, b []interface{}) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

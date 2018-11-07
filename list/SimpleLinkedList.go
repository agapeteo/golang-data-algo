package list

import (
	"errors"
)

var IndexOutOfBoundsError = errors.New("index out of bounds")
var EmptyListError = errors.New("list is empty")

type node struct {
	prev  *node
	next  *node
	value interface{}
}

type SimpleLinkedList struct {
	head *node
	tail *node
	size int
}

func (L *SimpleLinkedList) Size() int {
	return L.size
}

func (L *SimpleLinkedList) AddFirst(value interface{}) {
	newNode := &node{
		prev:  nil,
		next:  L.head,
		value: value,
	}

	if L.head == nil {
		L.head = newNode
		L.tail = newNode
	} else {
		L.head.prev = newNode
		newNode.next = L.head
		L.head = newNode
	}

	L.size++
}

func (L *SimpleLinkedList) AddLast(value interface{}) {
	newNode := &node{
		prev:  nil,
		next:  nil,
		value: value,
	}

	if L.tail == nil {
		L.head = newNode
		L.tail = newNode
	} else {
		curTail := L.tail
		newNode.prev = curTail
		curTail.next = newNode
		L.tail = newNode
	}

	L.size++
}

func (L *SimpleLinkedList) Append(value interface{}) {
	L.AddLast(value)
}

func (L *SimpleLinkedList) DeleteFirst() {
	if L.head == nil {
		return
	}

	oldHead := L.head
	newHead := oldHead.next
	L.head = newHead
	oldHead = nil

	L.size--
}

func (L *SimpleLinkedList) DeleteLast() {
	if L.tail == nil {
		return
	}

	oldTail := L.tail
	newTail := oldTail.prev
	newTail.next = nil
	L.tail = newTail
	oldTail = nil

	L.size--
}

func (L SimpleLinkedList) Elements() []interface{} {
	result := make([]interface{}, 0)

	if L.size == 0 {
		return result
	}

	currNode := L.head
	for currNode != nil {
		result = append(result, currNode.value)
		currNode = currNode.next
	}
	return result
}

func (L SimpleLinkedList) ReverseElements() []interface{} {
	result := make([]interface{}, 0)

	if L.size == 0 {
		return result
	}

	currNode := L.tail
	for currNode != nil {
		result = append(result, currNode.value)
		currNode = currNode.prev
	}
	return result
}

func (L *SimpleLinkedList) Get(index int) (value interface{}, error error) {
	if L.size == 0 {
		return nil, EmptyListError
	}
	if index < 0 || index >= L.size {
		return nil, IndexOutOfBoundsError
	}

	i := 0
	curNode := L.head

	for i < index {
		i++
		curNode = curNode.next
	}

	return curNode.value, nil
}

func (L *SimpleLinkedList) FirstIndexOf(element interface{}) int {
	if element == nil || L.size == 0 {
		return -1
	}

	curIndex := 0
	curNode := L.head

	for curNode != nil {
		if curNode.value == element {
			return curIndex
		}
		curIndex++
		curNode = curNode.next
	}

	return -1
}

func (L *SimpleLinkedList) InsertAfterIndex(index int, element interface{}) error {
	if index < 0 || index >= L.size {
		return IndexOutOfBoundsError
	}
	if L.size == 0 {
		return EmptyListError
	}

	curIndex := 0
	curNode := L.head

	for curIndex < index {
		curIndex++
		curNode = curNode.next
	}

	oldNext := curNode.next
	newNext := &node{curNode, oldNext, element}
	if oldNext != nil {
		oldNext.prev = newNext
	}
	curNode.next = newNext

	return nil
}

func (L *SimpleLinkedList) GetFirst() interface{} {
	if L.head == nil {
		return nil
	}
	return L.head.value
}

func (L *SimpleLinkedList) GetLast() interface{} {
	if L.tail == nil {
		return nil
	}
	return L.tail.value
}

package list

import (
	"log"
	"reflect"
)

func (l *SimpleLinkedList) SortInts() {
	if l.Size() < 2 {
		return
	}
	if reflect.TypeOf(l.head.value) != reflect.TypeOf(0) {
		log.Panicf("expected type is int, but got %v", reflect.TypeOf(l.head.value))
	}
	sortRange(l.head, l.tail)
}

func sortRange(startNode, endNode *node) {
	if startNode == endNode || endNode.next == startNode {
		return
	}
	partitionNode := partition(startNode, endNode)

	partitionPrev := partitionNode.prev
	if partitionPrev != nil {
		sortRange(startNode, partitionPrev)
	}

	partitionNext := partitionNode.next
	if partitionNext != nil {
		sortRange(partitionNext, endNode)
	}
}

func partition(startNode *node, endNode *node) *node {
	pivotValue := endNode.value.(int)
	slowNode := startNode
	fastNode := startNode

	for fastNode != endNode {
		if fastNode.value.(int) <= pivotValue {
			slowNode.value, fastNode.value = fastNode.value, slowNode.value
			slowNode = slowNode.next
		}
		fastNode = fastNode.next
	}
	slowNode.value, endNode.value = endNode.value, slowNode.value
	return slowNode
}

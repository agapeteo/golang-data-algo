package list

type LinkedListNode struct {
	value interface{}
	next  *LinkedListNode
}

func CreateLinkedListNode(value interface{}) *LinkedListNode {
	node := new(LinkedListNode)
	node.value = value
	return node
}

func (n *LinkedListNode) Append(value interface{}) {
	curNode := n
	for curNode.next != nil {
		curNode = curNode.next
	}
	curNode.next = CreateLinkedListNode(value)
}

func (n *LinkedListNode) Elements() []interface{} {
	result := make([]interface{}, 0)

	curNode := n
	for curNode != nil {
		result = append(result, curNode.value)
		curNode = curNode.next
	}
	return result
}

func (n *LinkedListNode) Size() int {
	count := 0
	curNode := n

	for curNode != nil {
		count++
		curNode = curNode.next
	}
	return count
}

func (n *LinkedListNode) RemoveAt(idx int) {
	size := n.Size()
	if size < 2 || idx < 1 || idx >= size {
		return
	}

	if size == 2 {
		n.next = nil
		return
	}

	prevNode := n
	curNode := n.next
	nextNode := curNode.next
	curIdx := 1

	for curIdx < idx && nextNode != nil {
		prevNode = prevNode.next
		curNode = curNode.next
		nextNode = nextNode.next
		curIdx++
	}
	prevNode.next = nextNode
}

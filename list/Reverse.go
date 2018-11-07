package list


func (L *SimpleLinkedList) Reverse(){
	if L.head == nil{
		return
	}

	if L.size == 2 {
		L.tail.next = L.head
		L.tail.prev = nil

		L.head.prev = L.tail
		L.head.next = nil

		swapHeadAndTail(L)
		return
	}

	var prevNode = L.head
	var curNode = prevNode.next
	var nextNode= curNode.next

	swapHeadAndTail(L)

	for curNode != nil {
		if prevNode.prev == nil{
			prevNode.next = nil
		}
		prevNode.prev = curNode

		curNode.next = prevNode
		curNode.prev = nextNode

		var nextNextNode *node
		if nextNode != nil{
			nextNextNode = nextNode.next
		}

		prevNode = curNode
		curNode = nextNode
		nextNode = nextNextNode
	}

}
func swapHeadAndTail(L *SimpleLinkedList) {
	var tmpHead = L.head
	L.head = L.tail
	L.tail = tmpHead
}

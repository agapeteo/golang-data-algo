package list

func (n *LinkedListNode) Deduplicate() {
	uniq := make(map[interface{}]struct{})
	curNode := n
	var prevNode *LinkedListNode

	for curNode != nil {
		_, exist := uniq[curNode.value]
		if exist {
			prevNode.next = curNode.next
		} else {
			prevNode = curNode
			uniq[curNode.value] = struct{}{}
		}
		curNode = curNode.next
	}
}

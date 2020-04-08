package trie

type node struct {
	nodes map[rune]*node
	last  bool
}

func newNode() *node {
	return &node{nodes: make(map[rune]*node), last: false}
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{root: newNode()}
}

func (t *Trie) Add(str string) {
	if len(str) == 0 {
		return
	}
	runes := []rune(str)
	addToNode(t.root, runes, 0)
}

func addToNode(n *node, runes []rune, idx int) {
	if idx >= len(runes) {
		return
	}
	r := runes[idx]

	var nextNode *node
	var ok bool
	if nextNode, ok = n.nodes[r]; !ok {
		nextNode = newNode()
		n.nodes[r] = nextNode
	}
	if idx == len(runes)-1 {
		nextNode.last = true
	}
	addToNode(nextNode, runes, idx+1)
}

func (t *Trie) Contains(str string) bool {
	if len(str) == 0 {
		return false
	}
	runes := []rune(str)
	return containsFromNode(t.root, runes, 0)
}

func containsFromNode(n *node, runes []rune, idx int) bool {
	if idx >= len(runes) {
		return false
	}
	r := runes[idx]

	var nextNode *node
	var ok bool
	if nextNode, ok = n.nodes[r]; !ok {
		return false
	}
	if idx == len(runes)-1 && nextNode.last {
		return true
	}
	return containsFromNode(nextNode, runes, idx+1)
}

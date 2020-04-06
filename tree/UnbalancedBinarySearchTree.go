package tree

type node struct {
	value int
	left  *node
	right *node
}

func newNode(value int) *node {
	return &node{value, nil, nil}
}

type UnbalancedBinarySearchTree struct {
	root *node
}

func (t *UnbalancedBinarySearchTree) Add(newValue int) {
	if t.root == nil {
		t.root = newNode(newValue)
		return
	}

	addToNode(t.root, newValue)
}

func FromSorted(s []int) UnbalancedBinarySearchTree {
	root := fromSortedSlice(s, 0, len(s))

	return UnbalancedBinarySearchTree{root}
}

func fromSortedSlice(s []int, start, end int) *node {
	if start >= end {
		return nil
	}
	mid := start + (end-start)/2
	value := s[mid]
	left := fromSortedSlice(s, start, mid)
	right := fromSortedSlice(s, mid+1, end)

	return &node{value, left, right}
}

func addToNode(n *node, newValue int) {
	if newValue <= n.value {
		if n.left == nil {
			n.left = newNode(newValue)
		} else {
			addToNode(n.left, newValue)
		}
	} else {
		if n.right == nil {
			n.right = newNode(newValue)
		} else {
			addToNode(n.right, newValue)
		}
	}
}

func (t *UnbalancedBinarySearchTree) OrderedElements() []int {
	if t.root == nil {
		return make([]int, 0)
	}

	var result = addToSlice(t.root, make([]int, 0))

	return result
}

func addToSlice(n *node, values []int) []int {
	if n == nil {
		return values
	}

	values = addToSlice(n.left, values)
	values = append(values, n.value)
	values = addToSlice(n.right, values)

	return values
}

func (t UnbalancedBinarySearchTree) Contains(value int) bool {
	if t.root == nil {
		return false
	}

	return containsInNode(t.root, value)
}

func containsInNode(n *node, value int) bool {
	if n.value == value {
		return true
	}

	if value < n.value {
		if n.left == nil {
			return false
		} else {
			return containsInNode(n.left, value)
		}
	} else {
		if n.right == nil {
			return false
		} else {
			return containsInNode(n.right, value)
		}
	}
}

func (t *UnbalancedBinarySearchTree) Depth() int {
	if t.root == nil {
		return 0
	}

	return nodeDepth(t.root, 0)
}

func nodeDepth(n *node, curDepth int) int {
	var leftDepth = curDepth
	var rightDepth = curDepth

	if n.left != nil {
		leftDepth = nodeDepth(n.left, curDepth+1)
	}

	if n.right != nil {
		rightDepth = nodeDepth(n.right, curDepth+1)
	}

	if leftDepth > rightDepth {
		return leftDepth
	} else {
		return rightDepth
	}
}
